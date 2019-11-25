package payloads

type ECProgram struct {
	ID           string      `json:"id"`
	Abbreviation string      `json:"abbreviation"`
	Type         interface{} `json:"type"`
	Program      interface{} `json:"program"`
	Description  string      `json:"description"`
	Home         bool        `json:"home,omitempty"`
}

type ECPrograms []ECProgram
