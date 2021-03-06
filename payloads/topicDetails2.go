package payloads

type Topics2 struct {
	TopicDetails struct {
		Type                int    `json:"type"`
		Ccm2ID              int    `json:"ccm2Id"`
		CftID               int    `json:"cftId"`
		Identifier          string `json:"identifier"`
		Title               string `json:"title"`
		PublicationDateLong int64  `json:"publicationDateLong"`
		CallIdentifier      string `json:"callIdentifier"`
		CallTitle           string `json:"callTitle"`
		Callccm2ID          int    `json:"callccm2Id"`
		WorkProgrammepart   struct {
			ID         int    `json:"id"`
			CcmID      int    `json:"ccm_id"`
			WpPart     string `json:"wp_part"`
			WpYear     string `json:"wp_year"`
			WpTitle    string `json:"wp_title"`
			WpWebsite  string `json:"wp_website"`
			WpDocument string `json:"wp_document"`
		} `json:"workProgrammepart"`
		FrameworkProgramme struct {
			ID           int    `json:"id"`
			Abbreviation string `json:"abbreviation"`
			Description  string `json:"description"`
		} `json:"frameworkProgramme"`
		ProgrammeDivision []struct {
			ID           int    `json:"id"`
			Abbreviation string `json:"abbreviation"`
			Description  string `json:"description"`
		} `json:"programmeDivision"`
		Tags    []string `json:"tags"`
		Sme     bool     `json:"sme"`
		Actions []struct {
			Status struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
				Description  string `json:"description"`
			} `json:"status"`
			Types               []string `json:"types"`
			PlannedOpeningDate  string   `json:"plannedOpeningDate"`
			SubmissionProcedure struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
				Description  string `json:"description"`
			} `json:"submissionProcedure"`
			DeadlineDates []string `json:"deadlineDates"`
		} `json:"actions"`
		LatestInfos            []interface{} `json:"latestInfos"`
		BudgetOverviewJSONItem struct {
			BudgetTopicActionMap struct {
				Num3248866 []struct {
					Action             string   `json:"action"`
					PlannedOpeningDate string   `json:"plannedOpeningDate"`
					DeadlineModel      string   `json:"deadlineModel"`
					DeadlineDates      []string `json:"deadlineDates"`
					BudgetYearMap      struct {
						Num2018 int `json:"2018"`
					} `json:"budgetYearMap"`
					BudgetTopicActionMap struct {
					} `json:"budgetTopicActionMap"`
				} `json:"3248866"`
			} `json:"budgetTopicActionMap"`
			BudgetYearsColumns []string `json:"budgetYearsColumns"`
		} `json:"budgetOverviewJSONItem"`
		Description        string        `json:"description"`
		Conditions         string        `json:"conditions"`
		SupportInfo        string        `json:"supportInfo"`
		SepTemplate        string        `json:"sepTemplate"`
		Links              []interface{} `json:"links"`
		AdditionalDossiers []struct {
			CallID           int64  `json:"callId"`
			ID               int    `json:"id"`
			Title            string `json:"title"`
			SectionID        int64  `json:"sectionId"`
			TranslationFiles []struct {
				Language string `json:"language"`
				MimeType string `json:"mimeType"`
				FileSize int    `json:"fileSize"`
				DocPath  string `json:"docPath"`
				ID       int    `json:"id"`
				FileName string `json:"fileName"`
				Common   bool   `json:"common"`
				Crc32    int64  `json:"crc32"`
			} `json:"translationFiles"`
		} `json:"additionalDossiers"`
		InfoPackDossiers    []interface{} `json:"infoPackDossiers"`
		CallDetailsJSONItem struct {
			AdditionalInfo string `json:"additionalInfo"`
			LatestInfos    []struct {
				ApprovalDate   string `json:"approvalDate"`
				LastChangeDate string `json:"lastChangeDate"`
				Content        string `json:"content"`
			} `json:"latestInfos"`
			HasForthcomingTopics bool `json:"hasForthcomingTopics"`
			HasOpenTopics        bool `json:"hasOpenTopics"`
			AllClosedTopics      bool `json:"allClosedTopics"`
		} `json:"callDetailsJSONItem"`
	} `json:"TopicDetails"`
}
