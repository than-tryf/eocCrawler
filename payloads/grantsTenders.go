package payloads

type GrantTenderObj struct {
	Type                   int     `json:"type"`
	Ccm2ID                 int     `json:"ccm2Id"`
	CftID                  int     `json:"cftId"`
	Identifier             string  `json:"identifier"`
	Title                  string  `json:"title"`
	PublicationDateLong    int64   `json:"publicationDateLong"`
	PlannedOpeningDateLong int64   `json:"plannedOpeningDateLong"`
	CallIdentifier         string  `json:"callIdentifier"`
	CallTitle              string  `json:"callTitle"`
	Callccm2ID             int     `json:"callccm2Id"`
	DeadlineDatesLong      []int64 `json:"deadlineDatesLong"`
	FrameworkProgramme     struct {
		ID           int    `json:"id"`
		Abbreviation string `json:"abbreviation"`
		Description  string `json:"description"`
	} `json:"frameworkProgramme"`
	ProgrammeDivision []struct {
		ID           int    `json:"id"`
		Abbreviation string `json:"abbreviation"`
		Description  string `json:"description"`
	} `json:"programmeDivision"`
	Status struct {
		ID           int    `json:"id"`
		Abbreviation string `json:"abbreviation"`
		Description  string `json:"description"`
	} `json:"status"`
	SumbissionProcedure struct {
		ID           int    `json:"id"`
		Abbreviation string `json:"abbreviation"`
		Description  string `json:"description"`
	} `json:"sumbissionProcedure"`
	TopicActions []struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	} `json:"topicActions"`
	Tags               []string      `json:"tags"`
	Flags              []string      `json:"flags,omitempty"`
	Sme                bool          `json:"sme"`
	Actions            []interface{} `json:"actions"`
	LatestInfos        []interface{} `json:"latestInfos"`
	Links              []interface{} `json:"links"`
	AdditionalDossiers []interface{} `json:"additionalDossiers"`
	InfoPackDossiers   []interface{} `json:"infoPackDossiers"`
	WorkProgrammepart  struct {
		ID         int    `json:"id"`
		CcmID      int    `json:"ccm_id"`
		WpPart     string `json:"wp_part"`
		WpYear     string `json:"wp_year"`
		WpTitle    string `json:"wp_title"`
		WpWebsite  string `json:"wp_website"`
		WpDocument string `json:"wp_document"`
	} `json:"workProgrammepart,omitempty"`
}

type GrantTenders struct {
	FundingData struct {
		GrantTenderObj []GrantTenderObj `json:"GrantTenderObj"`
	} `json:"fundingData"`
}
