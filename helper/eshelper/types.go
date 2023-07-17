package eshelper

import "encoding/json"

type ResponseBase struct {
	Index_       string `json:"_index"`
	Id_          string `json:"_id"`
	Version_     int    `json:"_version"`
	SeqNo_       int    `json:"_seq_no"`
	PrimaryTerm_ int    `json:"_primary_term"`
}

type IndexResult struct {
	ResponseBase
	Result  string `json:"result"`
	Shards_ Shards `json:"_shards"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type FetchDocResult struct {
	ResponseBase
	Found   bool            `json:"found"`
	Source_ json.RawMessage `json:"_source"`
}
