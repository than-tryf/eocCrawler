package payloads

import "github.com/jinzhu/gorm"

type ECProgram struct {
	gorm.Model
	ID           string `json:"id" gorm:"primary_key;column:programid"`
	Abbreviation string `json:"abbreviation"`
	Type         string `json:"type"`
	Program      string `json:"program"`
	Description  string `json:"description"`
	Home         bool   `json:"home,omitempty"`
}

type ECPrograms []ECProgram
