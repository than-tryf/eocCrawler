package main

import (
	"encoding/json"
	"eocCrawler/payloads"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	sources := []string{
		"https://eacea.ec.europa.eu/erasmus-plus/rss_en.xml",
		"https://eacea.ec.europa.eu/creative-europe/rss_en.xml",
		"https://eacea.ec.europa.eu/europe-for-citizens/rss_en.xml",
		"https://eacea.ec.europa.eu/eu-aid-volunteers/rss_en.xml_en",
	}

	wg.Add(len(sources))

	for _, source := range sources {
		go func(source string) {
			resp, _ := http.Get(source)
			rssObject := payloads.Rss{}
			xml.NewDecoder(resp.Body).Decode(&rssObject)

			for _, item := range rssObject.Channel.Item {
				if strings.Contains(item.Title, "Call for proposal") {
					if strings.Contains(item.Description, "<a href=") {
						document, _ := goquery.NewDocumentFromReader(strings.NewReader(item.Description))
						link, _ := document.Find("a").Attr("href")
						fmt.Printf("%v: %v: %v\n", rssObject.Channel.Title, item.Title, link)
					} else {
						fmt.Printf("%v: %v: %v\n", rssObject.Channel.Title, item.Title, "X")

					}

				}
			}

			wg.Done()
		}(source)
	}

	wg.Wait()

}
*/

const (
	CLOSED      = "Closed"
	FORTHCOMING = "Forthcoming"
	OPEN        = "Open"

	TENDERS_URL = "https://ec.europa.eu/info/funding-tenders/opportunities/data/referenceData/grantsTenders.json"
	TOPIC_URL   = "https://ec.europa.eu/info/funding-tenders/opportunities/data/topicDetails/"
)

func main() {
	start := time.Now()
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(2)

	resp, _ := http.Get(TENDERS_URL)

	callsData := payloads.GrantTenders{}
	_ = json.NewDecoder(resp.Body).Decode(&callsData)
	for _, call := range callsData.FundingData.GrantTenderObj {
		wg.Add(1)
		/*go func(callIndex int){
			if(call.Status.Description=="Forthcoming") {
				fmt.Printf("[%v] : %v\n\n", callIndex, call)
			}
			wg.Done()
		}(callIndex)*/

		go getCallTopics(call)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
	wg.Wait()
}

func getCallTopics(callTender payloads.GrantTenderObj) {
	if (callTender.Status.Description == OPEN || callTender.Status.Description == FORTHCOMING) && (callTender.Identifier) {

		//topicCall := strings.ToLower(callTender.Identifier)
		//resp, _ := http.Get(TOPIC_URL+topicCall+".json")
		//topicData := payloads.Topics{}
		//fmt.Println(topicCall)
		//_ = json.NewDecoder(resp.Body).Decode(&topicData)
		//fmt.Printf("************\n%v\n\n%v\n**********************\n", callTender, topicData)
		fmt.Printf("************\n%v\n***********\n", callTender)

	}
	wg.Done()
}
