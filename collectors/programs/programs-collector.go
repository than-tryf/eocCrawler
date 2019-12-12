package main

import (
	"encoding/json"
	"eocCrawler/payloads"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

const (
	PROGRAMS_URL = "https://ec.europa.eu/info/funding-tenders/opportunities/portal/assets/data/sedia-programmes.json"
)

func main() {

	db, err := gorm.Open("postgres", "host=10.16.3.61 port=31999 user=enedi_db_user dbname=pcod_dev password=j+qy5heKg@$-V6tv sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Printf("%v", err)
	}

	resp, _ := http.Get(PROGRAMS_URL)

	callPrograms := payloads.ECPrograms{}
	_ = json.NewDecoder(resp.Body).Decode(&callPrograms)

	// Migrate the schema
	//db.AutoMigrate(&payloads.ECProgram{})
	tbExists := db.HasTable(&payloads.ECProgram{})
	fmt.Printf("Not found creating: %v\n", tbExists)

	if !tbExists {
		fmt.Printf("CREATING: %v\n", tbExists)

		db.CreateTable(&payloads.ECProgram{})
		tbExists = db.HasTable(&payloads.ECProgram{})
		fmt.Printf("Should have now: %v\n", tbExists)
	}

	for _, callProgram := range callPrograms {
		db.Create(&callProgram)
	}

}
