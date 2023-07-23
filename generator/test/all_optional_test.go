package test

var sampleAllOptional AllOptional
var sampleAllOptionalZero = AllOptional{}

func wrapSliceEscape[T any](v T) *[]T {
	p := []T{v, v}
	return &p
}

func init() {
	sampleAllOptional = AllOptional{
		Agg:          escape(sampleAll.Agg),
		Blob:         wrapSliceEscape(sampleAll.Blob),
		Bool:         wrapSliceEscape(sampleAll.Bool),
		Byte:         wrapSliceEscape(sampleAll.Byte),
		Comp:         wrapSliceEscape(sampleAll.Comp),
		ConstantKwd:  wrapSliceEscape(sampleAll.ConstantKwd),
		Date:         wrapSliceEscape(AllOptionalDateDate(sampleAll.Date)),
		DateNano:     wrapSliceEscape(AllOptionalDateNanoDate(sampleAll.DateNano)),
		DateRange:    wrapSliceEscape(sampleAll.DateRange),
		DenseVector:  escape(sampleAll.DenseVector),
		Double:       wrapSliceEscape(sampleAll.Double),
		DoubleRange:  wrapSliceEscape(sampleAll.DoubleRange),
		Flattened:    wrapSliceEscape(sampleAll.Flattened),
		Float:        wrapSliceEscape(sampleAll.Float),
		FloatRange:   wrapSliceEscape(sampleAll.FloatRange),
		Geopoint:     wrapSliceEscape(sampleAll.Geopoint),
		Geoshape:     wrapSliceEscape(sampleAll.Geoshape),
		HalfFloat:    wrapSliceEscape(sampleAll.HalfFloat),
		Histogram:    escape(sampleAll.Histogram),
		Integer:      wrapSliceEscape(sampleAll.Integer),
		IntegerRange: wrapSliceEscape(sampleAll.IntegerRange),
		IpAddr:       wrapSliceEscape(sampleAll.IpAddr),
		IpRange:      wrapSliceEscape(sampleAll.IpRange),
		Join:         escape(sampleAll.Join),
		Kwd:          wrapSliceEscape(sampleAll.Kwd),
		Long:         wrapSliceEscape(sampleAll.Long),
		LongRange:    wrapSliceEscape(sampleAll.LongRange),
		// Nested and Object is omitted since it is tested in conversion and dynamic.
		Point:           wrapSliceEscape(sampleAll.Point),
		Query:           escape(sampleAll.Query),
		RankFeature:     escape(sampleAll.RankFeature),
		RankFeatures:    escape([]map[string]float64{{"foo": 12.3}, {"bar": 25.6}}),
		ScaledFloat:     wrapSliceEscape(sampleAll.ScaledFloat),
		SearchAsYouType: wrapSliceEscape(sampleAll.SearchAsYouType),
		Shape:           wrapSliceEscape(sampleAll.Shape),
		Short:           wrapSliceEscape(sampleAll.Short),
		Text:            wrapSliceEscape(sampleAll.Text),
		TextWTokenCount: wrapSliceEscape(sampleAll.TextWTokenCount),
		UnsignedLong:    wrapSliceEscape(sampleAll.UnsignedLong),
		Version:         wrapSliceEscape(sampleAll.Version),
		Wildcard:        wrapSliceEscape(sampleAll.Wildcard),
	}
}
