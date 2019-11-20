package main

import (
	"encoding/json"
	"eocCrawler/payloads"
	"eocCrawler/payloads/eoc"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	CLOSED      = "Closed"
	FORTHCOMING = "Forthcoming"
	OPEN        = "Open"

	TENDERS_URL = "https://ec.europa.eu/info/funding-tenders/opportunities/data/referenceData/grantsTenders.json"
	TOPIC_URL   = "https://ec.europa.eu/info/funding-tenders/opportunities/data/topicDetails/"

	GRANT  = 1
	TENDER = 0

	KAFKA_URL   = "10.16.3.22"
	KAFKA_TOPIC = "eocCalls"
)

var p *kafka.Producer

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

	p, _ = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": KAFKA_URL})

	defer p.Close()

	resp, _ := http.Get(TENDERS_URL)

	callsData := payloads.GrantTenders{}
	_ = json.NewDecoder(resp.Body).Decode(&callsData)
	for _, call := range callsData.FundingData.GrantTenderObj {
		wg.Add(1)
		go getCallTopics(call, &wg)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	wg.Wait()
}

func getCallTopics(callTender payloads.GrantTenderObj, wg *sync.WaitGroup) {
	if callTender.Status.Description == OPEN || callTender.Status.Description == FORTHCOMING {
		call4proposal := eoc.Call4Proposal{}
		var bts []byte
		if callTender.Type == GRANT {
			topicCall := strings.ToLower(callTender.Identifier)
			resp, err := http.Get(TOPIC_URL + topicCall + ".json")
			topicData := payloads.Topics{}
			call4proposal.Grant = callTender
			call4proposal.TopicInfo = payloads.TopicDetails{}
			if err == nil {
				_ = json.NewDecoder(resp.Body).Decode(&topicData)
				call4proposal.TopicInfo = topicData.TopicDetails
			}
			bts, _ = json.Marshal(call4proposal)
		} else {
			call4proposal.Grant = callTender
			call4proposal.TopicInfo = payloads.TopicDetails{}
			bts, _ = json.Marshal(call4proposal)
		}
		topic := KAFKA_TOPIC
		//callBodyBytes := new(bytes.Buffer)
		_ = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value:     bts,
			Timestamp: time.Time{},
		}, nil)

	}
	wg.Done()
}

/*func pushMessage(proposal chan eoc.Call4Proposal) {

}
*/
