package main

import (
	"encoding/json"
	"eocCrawler/payloads"
	"eocCrawler/payloads/eoc"
	"fmt"
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
)

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

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
		var bytes []byte
		if callTender.Type == GRANT {
			topicCall := strings.ToLower(callTender.Identifier)
			resp, _ := http.Get(TOPIC_URL + topicCall + ".json")
			topicData := payloads.Topics{}
			_ = json.NewDecoder(resp.Body).Decode(&topicData)
			call4proposal.Grant = callTender
			call4proposal.TopicInfo = topicData.TopicDetails
			bytes, _ = json.Marshal(call4proposal)
		} else {
			call4proposal.Grant = callTender
			call4proposal.TopicInfo = payloads.TopicDetails{}
			bytes, _ = json.Marshal(call4proposal)
		}
		fmt.Printf("%v\n", string(bytes))
	}
	wg.Done()
}
