package test

import (
	fielddatatype "github.com/ngicks/estype/fielddatatype"
	estime "github.com/ngicks/estype/fielddatatype/estime"
	builtin "github.com/ngicks/estype/fielddatatype/estime/builtin"
	gentypehelper "github.com/ngicks/estype/gentypehelper"
	elastic "github.com/ngicks/und/elastic"
	undefinedable "github.com/ngicks/und/undefinedable"
	"net/netip"
	"time"
)

type All struct {
	Agg             fielddatatype.AggregateMetricDouble  `json:"agg"`
	Alias           *struct{}                            `json:"alias,omitempty"`
	Blob            []byte                               `json:"blob"`
	Bool            fielddatatype.Boolean                `json:"bool"`
	Byte            int8                                 `json:"byte"`
	Comp            string                               `json:"comp"`
	ConstantKwd     string                               `json:"constant_kwd"`
	Date            AllDateDate                          `json:"date"`
	DateNano        AllDateNanoDate                      `json:"dateNano"`
	DateRange       fielddatatype.Range[builtin.Default] `json:"date_range"`
	DenseVector     [3]float64                           `json:"dense_vector"`
	Double          float64                              `json:"double"`
	DoubleRange     fielddatatype.Range[float64]         `json:"double_range"`
	Flattened       map[string]any                       `json:"flattened"`
	Float           float32                              `json:"float"`
	FloatRange      fielddatatype.Range[float32]         `json:"float_range"`
	Geopoint        fielddatatype.GeoPoint               `json:"geopoint"`
	Geoshape        fielddatatype.GeoShape               `json:"geoshape"`
	HalfFloat       float32                              `json:"half_float"`
	Histogram       fielddatatype.Histogram              `json:"histogram"`
	Integer         int32                                `json:"integer"`
	IntegerRange    fielddatatype.Range[int32]           `json:"integer_range"`
	IpAddr          netip.Addr                           `json:"ip_addr"`
	IpRange         fielddatatype.Range[netip.Addr]      `json:"ip_range"`
	Join            map[string]any                       `json:"join"`
	Kwd             string                               `json:"kwd"`
	Long            int64                                `json:"long"`
	LongRange       fielddatatype.Range[int64]           `json:"long_range"`
	Nested          AllNestedObject                      `json:"nested"`
	Object          AllObjectObject                      `json:"object"`
	Point           map[string]any                       `json:"point"`
	Query           map[string]any                       `json:"query"`
	RankFeature     float64                              `json:"rank_feature"`
	RankFeatures    map[string]float64                   `json:"rank_features"`
	ScaledFloat     float64                              `json:"scaled_float"`
	SearchAsYouType string                               `json:"search_as_you_type"`
	Shape           fielddatatype.GeoShape               `json:"shape"`
	Short           int16                                `json:"short"`
	Text            string                               `json:"text"`
	TextWTokenCount string                               `json:"text_w_token_count"`
	UnsignedLong    uint64                               `json:"unsigned_long"`
	Version         string                               `json:"version"`
	Wildcard        string                               `json:"wildcard"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d All) ToRaw() AllRaw {
	return AllRaw{
		Agg:             gentypehelper.MapSingleValueToElastic[fielddatatype.AggregateMetricDouble](d.Agg),
		Blob:            gentypehelper.MapSingleValueToElastic[[]byte](d.Blob),
		Bool:            gentypehelper.MapSingleValueToElastic[fielddatatype.Boolean](d.Bool),
		Byte:            gentypehelper.MapSingleValueToElastic[int8](d.Byte),
		Comp:            gentypehelper.MapSingleValueToElastic[string](d.Comp),
		ConstantKwd:     gentypehelper.MapSingleValueToElastic[string](d.ConstantKwd),
		Date:            gentypehelper.MapSingleValueToElastic[AllDateDate](d.Date),
		DateNano:        gentypehelper.MapSingleValueToElastic[AllDateNanoDate](d.DateNano),
		DateRange:       gentypehelper.MapSingleValueToElastic[fielddatatype.Range[builtin.Default]](d.DateRange),
		DenseVector:     undefinedable.Defined(d.DenseVector),
		Double:          gentypehelper.MapSingleValueToElastic[float64](d.Double),
		DoubleRange:     gentypehelper.MapSingleValueToElastic[fielddatatype.Range[float64]](d.DoubleRange),
		Flattened:       gentypehelper.MapSingleValueToElastic[map[string]any](d.Flattened),
		Float:           gentypehelper.MapSingleValueToElastic[float32](d.Float),
		FloatRange:      gentypehelper.MapSingleValueToElastic[fielddatatype.Range[float32]](d.FloatRange),
		Geopoint:        gentypehelper.MapSingleValueToElastic[fielddatatype.GeoPoint](d.Geopoint),
		Geoshape:        gentypehelper.MapSingleValueToElastic[fielddatatype.GeoShape](d.Geoshape),
		HalfFloat:       gentypehelper.MapSingleValueToElastic[float32](d.HalfFloat),
		Histogram:       gentypehelper.MapSingleValueToElastic[fielddatatype.Histogram](d.Histogram),
		Integer:         gentypehelper.MapSingleValueToElastic[int32](d.Integer),
		IntegerRange:    gentypehelper.MapSingleValueToElastic[fielddatatype.Range[int32]](d.IntegerRange),
		IpAddr:          gentypehelper.MapSingleValueToElastic[netip.Addr](d.IpAddr),
		IpRange:         gentypehelper.MapSingleValueToElastic[fielddatatype.Range[netip.Addr]](d.IpRange),
		Join:            gentypehelper.MapSingleValueToElastic[map[string]any](d.Join),
		Kwd:             gentypehelper.MapSingleValueToElastic[string](d.Kwd),
		Long:            gentypehelper.MapSingleValueToElastic[int64](d.Long),
		LongRange:       gentypehelper.MapSingleValueToElastic[fielddatatype.Range[int64]](d.LongRange),
		Nested:          gentypehelper.MapPlainToRawElastic[AllNestedObjectRaw](d.Nested),
		Object:          gentypehelper.MapPlainToRawElastic[AllObjectObjectRaw](d.Object),
		Point:           gentypehelper.MapSingleValueToElastic[map[string]any](d.Point),
		Query:           gentypehelper.MapSingleValueToElastic[map[string]any](d.Query),
		RankFeature:     gentypehelper.MapSingleValueToElastic[float64](d.RankFeature),
		RankFeatures:    gentypehelper.MapSingleValueToElastic[map[string]float64](d.RankFeatures),
		ScaledFloat:     gentypehelper.MapSingleValueToElastic[float64](d.ScaledFloat),
		SearchAsYouType: gentypehelper.MapSingleValueToElastic[string](d.SearchAsYouType),
		Shape:           gentypehelper.MapSingleValueToElastic[fielddatatype.GeoShape](d.Shape),
		Short:           gentypehelper.MapSingleValueToElastic[int16](d.Short),
		Text:            gentypehelper.MapSingleValueToElastic[string](d.Text),
		TextWTokenCount: gentypehelper.MapSingleValueToElastic[string](d.TextWTokenCount),
		UnsignedLong:    gentypehelper.MapSingleValueToElastic[uint64](d.UnsignedLong),
		Version:         gentypehelper.MapSingleValueToElastic[string](d.Version),
		Wildcard:        gentypehelper.MapSingleValueToElastic[string](d.Wildcard),
	}
}

type AllRaw struct {
	Agg             elastic.Elastic[fielddatatype.AggregateMetricDouble]  `json:"agg"`
	Alias           elastic.Elastic[*struct{}]                            `json:"alias"`
	Blob            elastic.Elastic[[]byte]                               `json:"blob"`
	Bool            elastic.Elastic[fielddatatype.Boolean]                `json:"bool"`
	Byte            elastic.Elastic[int8]                                 `json:"byte"`
	Comp            elastic.Elastic[string]                               `json:"comp"`
	ConstantKwd     elastic.Elastic[string]                               `json:"constant_kwd"`
	Date            elastic.Elastic[AllDateDate]                          `json:"date"`
	DateNano        elastic.Elastic[AllDateNanoDate]                      `json:"dateNano"`
	DateRange       elastic.Elastic[fielddatatype.Range[builtin.Default]] `json:"date_range"`
	DenseVector     undefinedable.Undefinedable[[3]float64]               `json:"dense_vector"`
	Double          elastic.Elastic[float64]                              `json:"double"`
	DoubleRange     elastic.Elastic[fielddatatype.Range[float64]]         `json:"double_range"`
	Flattened       elastic.Elastic[map[string]any]                       `json:"flattened"`
	Float           elastic.Elastic[float32]                              `json:"float"`
	FloatRange      elastic.Elastic[fielddatatype.Range[float32]]         `json:"float_range"`
	Geopoint        elastic.Elastic[fielddatatype.GeoPoint]               `json:"geopoint"`
	Geoshape        elastic.Elastic[fielddatatype.GeoShape]               `json:"geoshape"`
	HalfFloat       elastic.Elastic[float32]                              `json:"half_float"`
	Histogram       elastic.Elastic[fielddatatype.Histogram]              `json:"histogram"`
	Integer         elastic.Elastic[int32]                                `json:"integer"`
	IntegerRange    elastic.Elastic[fielddatatype.Range[int32]]           `json:"integer_range"`
	IpAddr          elastic.Elastic[netip.Addr]                           `json:"ip_addr"`
	IpRange         elastic.Elastic[fielddatatype.Range[netip.Addr]]      `json:"ip_range"`
	Join            elastic.Elastic[map[string]any]                       `json:"join"`
	Kwd             elastic.Elastic[string]                               `json:"kwd"`
	Long            elastic.Elastic[int64]                                `json:"long"`
	LongRange       elastic.Elastic[fielddatatype.Range[int64]]           `json:"long_range"`
	Nested          elastic.Elastic[AllNestedObjectRaw]                   `json:"nested"`
	Object          elastic.Elastic[AllObjectObjectRaw]                   `json:"object"`
	Point           elastic.Elastic[map[string]any]                       `json:"point"`
	Query           elastic.Elastic[map[string]any]                       `json:"query"`
	RankFeature     elastic.Elastic[float64]                              `json:"rank_feature"`
	RankFeatures    elastic.Elastic[map[string]float64]                   `json:"rank_features"`
	ScaledFloat     elastic.Elastic[float64]                              `json:"scaled_float"`
	SearchAsYouType elastic.Elastic[string]                               `json:"search_as_you_type"`
	Shape           elastic.Elastic[fielddatatype.GeoShape]               `json:"shape"`
	Short           elastic.Elastic[int16]                                `json:"short"`
	Text            elastic.Elastic[string]                               `json:"text"`
	TextWTokenCount elastic.Elastic[string]                               `json:"text_w_token_count"`
	UnsignedLong    elastic.Elastic[uint64]                               `json:"unsigned_long"`
	Version         elastic.Elastic[string]                               `json:"version"`
	Wildcard        elastic.Elastic[string]                               `json:"wildcard"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllRaw) ToPlain() All {
	return All{
		Agg:             d.Agg.ValueSingle(),
		Blob:            d.Blob.ValueSingle(),
		Bool:            d.Bool.ValueSingle(),
		Byte:            d.Byte.ValueSingle(),
		Comp:            d.Comp.ValueSingle(),
		ConstantKwd:     d.ConstantKwd.ValueSingle(),
		Date:            d.Date.ValueSingle(),
		DateNano:        d.DateNano.ValueSingle(),
		DateRange:       d.DateRange.ValueSingle(),
		DenseVector:     d.DenseVector.Value(),
		Double:          d.Double.ValueSingle(),
		DoubleRange:     d.DoubleRange.ValueSingle(),
		Flattened:       d.Flattened.ValueSingle(),
		Float:           d.Float.ValueSingle(),
		FloatRange:      d.FloatRange.ValueSingle(),
		Geopoint:        d.Geopoint.ValueSingle(),
		Geoshape:        d.Geoshape.ValueSingle(),
		HalfFloat:       d.HalfFloat.ValueSingle(),
		Histogram:       d.Histogram.ValueSingle(),
		Integer:         d.Integer.ValueSingle(),
		IntegerRange:    d.IntegerRange.ValueSingle(),
		IpAddr:          d.IpAddr.ValueSingle(),
		IpRange:         d.IpRange.ValueSingle(),
		Join:            d.Join.ValueSingle(),
		Kwd:             d.Kwd.ValueSingle(),
		Long:            d.Long.ValueSingle(),
		LongRange:       d.LongRange.ValueSingle(),
		Nested:          gentypehelper.MapElasticToPlainSingle[AllNestedObject](d.Nested),
		Object:          gentypehelper.MapElasticToPlainSingle[AllObjectObject](d.Object),
		Point:           d.Point.ValueSingle(),
		Query:           d.Query.ValueSingle(),
		RankFeature:     d.RankFeature.ValueSingle(),
		RankFeatures:    d.RankFeatures.ValueSingle(),
		ScaledFloat:     d.ScaledFloat.ValueSingle(),
		SearchAsYouType: d.SearchAsYouType.ValueSingle(),
		Shape:           d.Shape.ValueSingle(),
		Short:           d.Short.ValueSingle(),
		Text:            d.Text.ValueSingle(),
		TextWTokenCount: d.TextWTokenCount.ValueSingle(),
		UnsignedLong:    d.UnsignedLong.ValueSingle(),
		Version:         d.Version.ValueSingle(),
		Wildcard:        d.Wildcard.ValueSingle(),
	}
}

// AllDateDate represents the date or the date_nanos mapping field type.
// It implements json.Unmarshaler so that it can be directly unmarshaled from
// all possible formats specified in corresponding `format` field.
//
// Allowed formats are:
//
//   - 2006-01-02 15:04:05
//   - 2006-01-02
//   - int as epoch_millis
//
// It also implements json.Marshaler. It will be marshaled into
// string formatted in 2006-01-02 15:04:05 layout
type AllDateDate time.Time

var parserAllDateDate = estime.FromGoTimeLayoutUnsafe(
	[]string{
		"2006-01-02 15:04:05",
		"2006-01-02",
	},
	"epoch_millis",
)

// String implements fmt.Stringer
func (t AllDateDate) String() string {
	return parserAllDateDate.FormatString(time.Time(t), 0)
}

// MarshalJSON implements json.Marshaler
func (t AllDateDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (t *AllDateDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	tt, err := parserAllDateDate.ParseJson(data)
	if err != nil {
		return err
	}
	*t = AllDateDate(tt)
	return nil
}

// AllDateNanoDate represents the date or the date_nanos mapping field type.
// It implements json.Unmarshaler so that it can be directly unmarshaled from
// all possible formats specified in corresponding `format` field.
//
// Allowed formats are:
//
//   - 2006-01-02T15:04:05.000000000Z0700
//   - 2006-01-02
//   - int as epoch_second
//
// It also implements json.Marshaler. It will be marshaled into
// string formatted in 2006-01-02T15:04:05.000000000Z0700 layout
type AllDateNanoDate time.Time

var parserAllDateNanoDate = estime.FromGoTimeLayoutUnsafe(
	[]string{
		"2006-01-02T15:04:05.000000000Z0700",
		"2006-01-02",
	},
	"epoch_second",
)

// String implements fmt.Stringer
func (t AllDateNanoDate) String() string {
	return parserAllDateNanoDate.FormatString(time.Time(t), 0)
}

// MarshalJSON implements json.Marshaler
func (t AllDateNanoDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (t *AllDateNanoDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	tt, err := parserAllDateNanoDate.ParseJson(data)
	if err != nil {
		return err
	}
	*t = AllDateNanoDate(tt)
	return nil
}

type AllNestedObject struct {
	Age  int32               `json:"age"`
	Name AllNestedNameObject `json:"name"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllNestedObject) ToRaw() AllNestedObjectRaw {
	return AllNestedObjectRaw{
		Age:  gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainToRawElastic[AllNestedNameObjectRaw](d.Name),
	}
}

type AllNestedObjectRaw struct {
	Age  elastic.Elastic[int32]                  `json:"age"`
	Name elastic.Elastic[AllNestedNameObjectRaw] `json:"name"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllNestedObjectRaw) ToPlain() AllNestedObject {
	return AllNestedObject{
		Age:  d.Age.ValueSingle(),
		Name: gentypehelper.MapElasticToPlainSingle[AllNestedNameObject](d.Name),
	}
}

type AllNestedNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllNestedNameObject) ToRaw() AllNestedNameObjectRaw {
	return AllNestedNameObjectRaw{
		First: gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:  gentypehelper.MapSingleValueToElastic[string](d.Last),
	}
}

type AllNestedNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllNestedNameObjectRaw) ToPlain() AllNestedNameObject {
	return AllNestedNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}

type AllObjectObject struct {
	Age  int32               `json:"age"`
	Name AllObjectNameObject `json:"name"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllObjectObject) ToRaw() AllObjectObjectRaw {
	return AllObjectObjectRaw{
		Age:  gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainToRawElastic[AllObjectNameObjectRaw](d.Name),
	}
}

type AllObjectObjectRaw struct {
	Age  elastic.Elastic[int32]                  `json:"age"`
	Name elastic.Elastic[AllObjectNameObjectRaw] `json:"name"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllObjectObjectRaw) ToPlain() AllObjectObject {
	return AllObjectObject{
		Age:  d.Age.ValueSingle(),
		Name: gentypehelper.MapElasticToPlainSingle[AllObjectNameObject](d.Name),
	}
}

type AllObjectNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllObjectNameObject) ToRaw() AllObjectNameObjectRaw {
	return AllObjectNameObjectRaw{
		First: gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:  gentypehelper.MapSingleValueToElastic[string](d.Last),
	}
}

type AllObjectNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllObjectNameObjectRaw) ToPlain() AllObjectNameObject {
	return AllObjectNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}
