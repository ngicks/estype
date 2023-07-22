package test

import (
	"github.com/ngicks/generic"
)

var sampleAllOptional AllOptional

func wrapSliceEscape[T any](v T) *[]T {
	p := []T{v}
	return &p
}

func init() {
	sampleAllOptional = AllOptional{
		Agg:          generic.Escape(sampleAll.Agg),
		Blob:         wrapSliceEscape(sampleAll.Blob),
		Bool:         wrapSliceEscape(sampleAll.Bool),
		Byte:         wrapSliceEscape(sampleAll.Byte),
		Comp:         wrapSliceEscape(sampleAll.Comp),
		ConstantKwd:  wrapSliceEscape(sampleAll.ConstantKwd),
		Date:         wrapSliceEscape(AllOptionalDateDate(sampleAll.Date)),
		DateNano:     wrapSliceEscape(AllOptionalDateNanoDate(sampleAll.DateNano)),
		DateRange:    wrapSliceEscape(sampleAll.DateRange),
		DenseVector:  generic.Escape(sampleAll.DenseVector),
		Double:       wrapSliceEscape(sampleAll.Double),
		DoubleRange:  wrapSliceEscape(sampleAll.DoubleRange),
		Flattened:    wrapSliceEscape(sampleAll.Flattened),
		Float:        wrapSliceEscape(sampleAll.Float),
		FloatRange:   wrapSliceEscape(sampleAll.FloatRange),
		Geopoint:     wrapSliceEscape(sampleAll.Geopoint),
		Geoshape:     wrapSliceEscape(sampleAll.Geoshape),
		HalfFloat:    wrapSliceEscape(sampleAll.HalfFloat),
		Histogram:    wrapSliceEscape(sampleAll.Histogram),
		Integer:      wrapSliceEscape(sampleAll.Integer),
		IntegerRange: wrapSliceEscape(sampleAll.IntegerRange),
		IpAddr:       wrapSliceEscape(sampleAll.IpAddr),
		IpRange:      wrapSliceEscape(sampleAll.IpRange),
		Join:         wrapSliceEscape(sampleAll.Join),
		Kwd:          wrapSliceEscape(sampleAll.Kwd),
		Long:         wrapSliceEscape(sampleAll.Long),
		LongRange:    wrapSliceEscape(sampleAll.LongRange),
		// Nested and Object is omitted since it is tested in conversion and dynamic.
		Point:           wrapSliceEscape(sampleAll.Point),
		Query:           wrapSliceEscape(sampleAll.Query),
		RankFeature:     wrapSliceEscape(sampleAll.RankFeature),
		RankFeatures:    wrapSliceEscape(sampleAll.RankFeatures),
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
