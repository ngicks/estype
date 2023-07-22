package fielddatatype

type Histogram struct {
	Values []float64 `json:"values"`
	Counts []int32   `json:"counts"`
}
