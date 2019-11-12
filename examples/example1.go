package main

import (
	"bytes"
	"encoding/json"
	"eocCrawler/payloads"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "10.16.3.22"})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	sources := []string{
		"https://ec.europa.eu/research/participants/portal/data/call/h2020/calls.json",
	}

	wg.Add(len(sources))

	for _, source := range sources {
		go func(source string) {

			resp, _ := http.Get(source)

			callsData := payloads.Call{}
			_ = json.NewDecoder(resp.Body).Decode(&callsData)
			fmt.Println(len(callsData.CallData.Calls))
			topic := "eocCalls"
			for _, call := range callsData.CallData.Calls {
				callBodyBytes := new(bytes.Buffer)
				json.NewEncoder(callBodyBytes).Encode(call)
				p.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{
						Topic:     &topic,
						Partition: kafka.PartitionAny,
					},
					Value:     callBodyBytes.Bytes(),
					Timestamp: time.Time{},
				}, nil)
			}

			p.Flush(15 * 1000)
			elapsed := time.Since(start)
			fmt.Println(elapsed)
			wg.Done()
		}(source)
	}

	wg.Wait()
}
