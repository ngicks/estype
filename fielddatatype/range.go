package fielddatatype

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/range.html
type Range[T comparable] struct {
	Gt  *T `json:"gt,omitempty"`
	Gte *T `json:"gte,omitempty"`
	Lt  *T `json:"lt,omitempty"`
	Lte *T `json:"lte,omitempty"`
}
