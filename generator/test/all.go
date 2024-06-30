package test

import (
	fielddatatype "github.com/ngicks/estype/fielddatatype"
	estime "github.com/ngicks/estype/fielddatatype/estime"
	builtin "github.com/ngicks/estype/fielddatatype/estime/builtin"
	gentypehelper "github.com/ngicks/estype/gentypehelper"
	sliceund "github.com/ngicks/und/sliceund"
	elastic "github.com/ngicks/und/sliceund/elastic"
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
		DenseVector:     sliceund.Defined(d.DenseVector),
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
	Agg             elastic.Elastic[fielddatatype.AggregateMetricDouble]  `json:"agg,omitempty"`
	Alias           elastic.Elastic[*struct{}]                            `json:"alias,omitempty"`
	Blob            elastic.Elastic[[]byte]                               `json:"blob,omitempty"`
	Bool            elastic.Elastic[fielddatatype.Boolean]                `json:"bool,omitempty"`
	Byte            elastic.Elastic[int8]                                 `json:"byte,omitempty"`
	Comp            elastic.Elastic[string]                               `json:"comp,omitempty"`
	ConstantKwd     elastic.Elastic[string]                               `json:"constant_kwd,omitempty"`
	Date            elastic.Elastic[AllDateDate]                          `json:"date,omitempty"`
	DateNano        elastic.Elastic[AllDateNanoDate]                      `json:"dateNano,omitempty"`
	DateRange       elastic.Elastic[fielddatatype.Range[builtin.Default]] `json:"date_range,omitempty"`
	DenseVector     sliceund.Und[[3]float64]                              `json:"dense_vector,omitempty"`
	Double          elastic.Elastic[float64]                              `json:"double,omitempty"`
	DoubleRange     elastic.Elastic[fielddatatype.Range[float64]]         `json:"double_range,omitempty"`
	Flattened       elastic.Elastic[map[string]any]                       `json:"flattened,omitempty"`
	Float           elastic.Elastic[float32]                              `json:"float,omitempty"`
	FloatRange      elastic.Elastic[fielddatatype.Range[float32]]         `json:"float_range,omitempty"`
	Geopoint        elastic.Elastic[fielddatatype.GeoPoint]               `json:"geopoint,omitempty"`
	Geoshape        elastic.Elastic[fielddatatype.GeoShape]               `json:"geoshape,omitempty"`
	HalfFloat       elastic.Elastic[float32]                              `json:"half_float,omitempty"`
	Histogram       elastic.Elastic[fielddatatype.Histogram]              `json:"histogram,omitempty"`
	Integer         elastic.Elastic[int32]                                `json:"integer,omitempty"`
	IntegerRange    elastic.Elastic[fielddatatype.Range[int32]]           `json:"integer_range,omitempty"`
	IpAddr          elastic.Elastic[netip.Addr]                           `json:"ip_addr,omitempty"`
	IpRange         elastic.Elastic[fielddatatype.Range[netip.Addr]]      `json:"ip_range,omitempty"`
	Join            elastic.Elastic[map[string]any]                       `json:"join,omitempty"`
	Kwd             elastic.Elastic[string]                               `json:"kwd,omitempty"`
	Long            elastic.Elastic[int64]                                `json:"long,omitempty"`
	LongRange       elastic.Elastic[fielddatatype.Range[int64]]           `json:"long_range,omitempty"`
	Nested          elastic.Elastic[AllNestedObjectRaw]                   `json:"nested,omitempty"`
	Object          elastic.Elastic[AllObjectObjectRaw]                   `json:"object,omitempty"`
	Point           elastic.Elastic[map[string]any]                       `json:"point,omitempty"`
	Query           elastic.Elastic[map[string]any]                       `json:"query,omitempty"`
	RankFeature     elastic.Elastic[float64]                              `json:"rank_feature,omitempty"`
	RankFeatures    elastic.Elastic[map[string]float64]                   `json:"rank_features,omitempty"`
	ScaledFloat     elastic.Elastic[float64]                              `json:"scaled_float,omitempty"`
	SearchAsYouType elastic.Elastic[string]                               `json:"search_as_you_type,omitempty"`
	Shape           elastic.Elastic[fielddatatype.GeoShape]               `json:"shape,omitempty"`
	Short           elastic.Elastic[int16]                                `json:"short,omitempty"`
	Text            elastic.Elastic[string]                               `json:"text,omitempty"`
	TextWTokenCount elastic.Elastic[string]                               `json:"text_w_token_count,omitempty"`
	UnsignedLong    elastic.Elastic[uint64]                               `json:"unsigned_long,omitempty"`
	Version         elastic.Elastic[string]                               `json:"version,omitempty"`
	Wildcard        elastic.Elastic[string]                               `json:"wildcard,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllRaw) ToPlain() All {
	return All{
		Agg:             d.Agg.Value(),
		Blob:            d.Blob.Value(),
		Bool:            d.Bool.Value(),
		Byte:            d.Byte.Value(),
		Comp:            d.Comp.Value(),
		ConstantKwd:     d.ConstantKwd.Value(),
		Date:            d.Date.Value(),
		DateNano:        d.DateNano.Value(),
		DateRange:       d.DateRange.Value(),
		DenseVector:     d.DenseVector.Value(),
		Double:          d.Double.Value(),
		DoubleRange:     d.DoubleRange.Value(),
		Flattened:       d.Flattened.Value(),
		Float:           d.Float.Value(),
		FloatRange:      d.FloatRange.Value(),
		Geopoint:        d.Geopoint.Value(),
		Geoshape:        d.Geoshape.Value(),
		HalfFloat:       d.HalfFloat.Value(),
		Histogram:       d.Histogram.Value(),
		Integer:         d.Integer.Value(),
		IntegerRange:    d.IntegerRange.Value(),
		IpAddr:          d.IpAddr.Value(),
		IpRange:         d.IpRange.Value(),
		Join:            d.Join.Value(),
		Kwd:             d.Kwd.Value(),
		Long:            d.Long.Value(),
		LongRange:       d.LongRange.Value(),
		Nested:          gentypehelper.MapElasticToPlainSingle[AllNestedObject](d.Nested),
		Object:          gentypehelper.MapElasticToPlainSingle[AllObjectObject](d.Object),
		Point:           d.Point.Value(),
		Query:           d.Query.Value(),
		RankFeature:     d.RankFeature.Value(),
		RankFeatures:    d.RankFeatures.Value(),
		ScaledFloat:     d.ScaledFloat.Value(),
		SearchAsYouType: d.SearchAsYouType.Value(),
		Shape:           d.Shape.Value(),
		Short:           d.Short.Value(),
		Text:            d.Text.Value(),
		TextWTokenCount: d.TextWTokenCount.Value(),
		UnsignedLong:    d.UnsignedLong.Value(),
		Version:         d.Version.Value(),
		Wildcard:        d.Wildcard.Value(),
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
// string formatted in 2006-01-02 15:04:05 layout.
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
// string formatted in 2006-01-02T15:04:05.000000000Z0700 layout.
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
	Age  elastic.Elastic[int32]                  `json:"age,omitempty"`
	Name elastic.Elastic[AllNestedNameObjectRaw] `json:"name,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllNestedObjectRaw) ToPlain() AllNestedObject {
	return AllNestedObject{
		Age:  d.Age.Value(),
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
	First elastic.Elastic[string] `json:"first,omitempty"`
	Last  elastic.Elastic[string] `json:"last,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllNestedNameObjectRaw) ToPlain() AllNestedNameObject {
	return AllNestedNameObject{
		First: d.First.Value(),
		Last:  d.Last.Value(),
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
	Age  elastic.Elastic[int32]                  `json:"age,omitempty"`
	Name elastic.Elastic[AllObjectNameObjectRaw] `json:"name,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllObjectObjectRaw) ToPlain() AllObjectObject {
	return AllObjectObject{
		Age:  d.Age.Value(),
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
	First elastic.Elastic[string] `json:"first,omitempty"`
	Last  elastic.Elastic[string] `json:"last,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllObjectNameObjectRaw) ToPlain() AllObjectNameObject {
	return AllObjectNameObject{
		First: d.First.Value(),
		Last:  d.Last.Value(),
	}
}
