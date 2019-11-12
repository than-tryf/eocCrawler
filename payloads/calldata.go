package payloads

type Call struct {
	CallData struct {
		Calls []struct {
			CallIdentifier struct {
				FileName string `json:"FileName"`
				CallID   string `json:"CallId"`
				Status   string `json:"Status"`
			} `json:"CallIdentifier"`
			Title                                  string        `json:"Title"`
			FrameworkProgramme                     string        `json:"FrameworkProgramme"`
			SpecificProgrammeLevel1Names           []interface{} `json:"SpecificProgrammeLevel1Names"`
			SpecificProgrammeLevel2Names           []interface{} `json:"SpecificProgrammeLevel2Names"`
			SpecificProgrammeNames                 []interface{} `json:"SpecificProgrammeNames"`
			MainSpecificProgrammeLevel1Name        string        `json:"MainSpecificProgrammeLevel1Name"`
			MainSpecificProgrammeLevel1Description string        `json:"MainSpecificProgrammeLevel1Description"`
			PublicationDate                        int64         `json:"PublicationDate"`
			DeadlineDate                           int64         `json:"DeadlineDate"`
			Deadlines                              []string      `json:"Deadlines"`
			CallDetails                            struct {
				AdditionalInfo string `json:"AdditionalInfo"`
				LatestInfo     struct {
				} `json:"LatestInfo"`
			} `json:"CallDetails"`
			Theme    []interface{} `json:"Theme"`
			Type     string        `json:"Type"`
			Category string        `json:"Category"`
		} `json:"Calls"`
	} `json:"callData"`
}
