package eoc

import "eocCrawler/payloads"

type Call4Proposal struct {
	Grant     payloads.GrantTenderObj `json:"grant"`
	TopicInfo payloads.TopicDetails   `json:"topicinfo"`
}
