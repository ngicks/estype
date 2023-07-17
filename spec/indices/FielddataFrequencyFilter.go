package indices

type FielddataFrequencyFilter struct {
	Max            float64 `json:"max"`
	Min            float64 `json:"min"`
	MinSegmentSize int     `json:"min_segment_size"`
}
