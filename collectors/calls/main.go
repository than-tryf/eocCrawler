package main

import (
	"encoding/json"
	"eocCrawler/payloads"
	"eocCrawler/payloads/eoc"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
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

	KAFKA_URL   = "platform.eoc.org.cy"
	KAFKA_PORT  = "31090"
	KAFKA_TOPIC = "eocCalls"

	NUM_WORKERS = 40
)

var p *kafka.Producer

func main() {
	var c4p [NUM_WORKERS]<-chan eoc.Call4Proposal
	calls := make(chan payloads.GrantTenderObj)
	done := make(chan interface{})
	start := time.Now()
	//p, _ = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": KAFKA_URL + ":" + KAFKA_PORT})
	//defer p.Close()
	resp, _ := http.Get(TENDERS_URL)
	callsData := payloads.GrantTenders{}
	err := json.NewDecoder(resp.Body).Decode(&callsData)
	http.Client.
	if err!=nil {
		log.Fatalf("Error decoding data: %v\n", err)
	}
	for wrk:=0; wrk < NUM_WORKERS; wrk++{
		c4p[wrk] = make(chan  eoc.Call4Proposal)
		c4p[wrk] = getCallTopicsFanOut(done, calls, &wg)
	}
	log.Println("Num of CPU: ", runtime.NumCPU())
	log.Println("Num of Goroutine: ", runtime.NumGoroutine())

	for _, call := range callsData.FundingData.GrantTenderObj{
		calls <- call
	}
	close(done)

	/*for _, call := range callsData.FundingData.GrantTenderObj {
		wg.Add(1)
		go getCallTopics(call, &wg)
	}*/

	/*wg.Add(NUM_WORKERS)
	for wrk:=1; wrk<=NUM_WORKERS; wrk++ {
		go getCallTopicsBuffered(calls, wrk)
	}

	f

	for _,call := range callsData.FundingData.GrantTenderObj{
		calls <- call
	}

	close(calls)*/
	wg.Wait()
	close(calls)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func getCallTopicsFanOut(done <-chan interface{}, calls <- chan payloads.GrantTenderObj, wg *sync.WaitGroup) <-chan eoc.Call4Proposal  {
	c4p := make(chan eoc.Call4Proposal)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(c4p)
		for {
			select {
			case <-done:
				fmt.Println("Stopping: ", &c4p)
				return
			case callTender:= <- calls:
				fmt.Println("Working: ", &c4p, " ", callTender.CallTitle)
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
							call4proposal.TopicInfo = topicData.TopicDetails
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
					c4p <- call4proposal

				}
			default:

			}
		}
	}()
	return c4p
}

func getCallTopics(callTender payloads.GrantTenderObj, wg *sync.WaitGroup) {
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
		}, nil)*/

		fmt.Printf("\n\n%v\n\n", string(bts))

	}
}


func getCallTopicsBuffered(calls chan payloads.GrantTenderObj, worker int) {
	defer wg.Done()
	for {
		callTender, ok := <- calls
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// Display we are starting the work.
		fmt.Printf("Worker: %d \n", worker)

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

			//fmt.Printf("%v", bts)

		}

		fmt.Printf("Worker: %d Completed \n", worker)

	}

}