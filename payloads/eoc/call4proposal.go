package eoc

import "eocCrawler/payloads"

type Call4Proposal struct {
	Grant     payloads.GrantTenderObj
	TopicInfo payloads.TopicDetails
}
