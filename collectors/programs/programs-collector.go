package main

import (
	"encoding/json"
	"eocCrawler/payloads"
	"fmt"
	"net/http"
)

const (
	PROGRAMS_URL = "https://ec.europa.eu/info/funding-tenders/opportunities/portal/assets/data/sedia-programmes.json"
)

func main() {
	resp, _ := http.Get(PROGRAMS_URL)

	callPrograms := payloads.ECPrograms{}
	_ = json.NewDecoder(resp.Body).Decode(&callPrograms)

	for _, callProgram := range callPrograms {
		fmt.Printf("%v\n", callProgram.Abbreviation)
	}

}
