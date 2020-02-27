package main

import (
	"encoding/json"
	"eocCrawler/payloads"
	"eocCrawler/payloads/eoc"
	"fmt"
	"log"
	"net/http"
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

	KAFKA_URL   = "platform.eoc.org.cy"
	KAFKA_PORT  = "31090"
	KAFKA_TOPIC = "eocCalls"

	NUM_WORKERS = 64
)

//var p *kafka.Producer

func main() {
	var c4p [NUM_WORKERS]chan eoc.Call4Proposal
	calls := make(chan payloads.GrantTenderObj)
	done := make(chan interface{})
	start := time.Now()
	resp, _ := http.Get(TENDERS_URL)
	callsData := payloads.GrantTenders{}
	err := json.NewDecoder(resp.Body).Decode(&callsData)
	_ = resp.Body.Close()
	if err!=nil {
		log.Fatalf("Error decoding data: %v\n", err)
	}
	for wrk:=0; wrk < NUM_WORKERS; wrk++{
		c4p[wrk] = make(chan  eoc.Call4Proposal)
		c4p[wrk] = concurrentJob(done, calls, &wg)
	}
	//numOfCalls:=0
	go func() {
		for n := range merge(c4p[:]...) {
			//fmt.Println(n)
			//numOfCalls++
			fmt.Println(n.Grant.Title)
		}
	}()
	for _, call := range callsData.FundingData.GrantTenderObj{
		calls <- call
	}


	close(done)
	close(calls)
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func concurrentJob(done <-chan interface{}, calls <- chan payloads.GrantTenderObj, wg *sync.WaitGroup) chan eoc.Call4Proposal {
	c := make(chan eoc.Call4Proposal)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case callTender := <-calls:
				{
					if callTender.Status.Description == OPEN || callTender.Status.Description == FORTHCOMING {
						call4proposal := eoc.Call4Proposal{}
						//var bts []byte
						if callTender.Type == GRANT {
							topicCall := strings.ToLower(callTender.Identifier)
							resp, err := http.Get(TOPIC_URL + topicCall + ".json")
							topicData := payloads.Topics{}
							call4proposal.Grant = callTender
							call4proposal.TopicInfo = payloads.TopicDetails{}
							if err == nil {
								_ = json.NewDecoder(resp.Body).Decode(&topicData)
								_ = resp.Body.Close()
								call4proposal.TopicInfo = topicData.TopicDetails
							} else {
								fmt.Println("PROBLEM WITH ID")
							}
							//bts, _ = json.Marshal(call4proposal)
						} else {
							call4proposal.Grant = callTender
							call4proposal.TopicInfo = payloads.TopicDetails{}
							//bts, _ = json.Marshal(call4proposal)
						}

						/*topic := KAFKA_TOPIC
						//callBodyBytes := new(bytes.Buffer)
						_ = p.Produce(&kafka.Message{
							TopicPartition: kafka.TopicPartition{
								Topic:     &topic,
								Partition: kafka.PartitionAny,
							},
							Value:     bts,
							Timestamp: time.Time{},
						}, nil)*/

						//fmt.Printf("\n\n%v\n\n", string(bts))
						c <- call4proposal

					}
				}
			}
		}
	}()
	return c
}

func merge(cs ...chan eoc.Call4Proposal) <-chan eoc.Call4Proposal {
	var wg sync.WaitGroup
	out := make(chan eoc.Call4Proposal)
	output := func(c <-chan eoc.Call4Proposal) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		fmt.Printf("Num of Channels: %v\n", len(cs))
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}


/*func getCallTopics(callTender payloads.GrantTenderObj, wg *sync.WaitGroup) {
	defer wg.Done()
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

		/*topic := KAFKA_TOPIC
		//callBodyBytes := new(bytes.Buffer)
		_ = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value:     bts,
			Timestamp: time.Time{},
		}, nil)

		fmt.Printf("\n\n%v\n\n", string(bts))

	}
}
*/
