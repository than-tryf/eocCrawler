package main

import (
	"sync"
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
//
//func main() {
//	start := time.Now()
//	runtime.GOMAXPROCS(runtime.NumCPU())
//
//	sources := []string{
//		"https://ec.europa.eu/research/participants/portal/data/call/h2020/calls.json",
//	}
//
//	wg.Add(len(sources))
//
//	for _, source := range sources {
//		go func(source string){
//
//			resp, _ := http.Get(source)
//
//			callsData := payloads.Call{}
//			_ = json.NewDecoder(resp.Body).Decode(&callsData)
//			fmt.Println(len(callsData.CallData.Calls))
//
//			elapsed := time.Since(start)
//			fmt.Println(elapsed)
//			wg.Done()
//		}(source)
//	}
//
//	wg.Wait()
//
//
//}
