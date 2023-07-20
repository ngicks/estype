package test

import (
	fielddatatype "github.com/ngicks/estype/fielddatatype"
	estime "github.com/ngicks/estype/fielddatatype/estime"
	builtin "github.com/ngicks/estype/fielddatatype/estime/builtin"
	elastic "github.com/ngicks/und/elastic"
	"net/netip"
	"time"
)

type All struct {
	Agg             fielddatatype.AggregateMetricDouble  `json:"agg"`
	Alias           any                                  `json:"alias,omitempty"`
	Blob            []byte                               `json:"blob"`
	Bool            fielddatatype.Boolean                `json:"bool"`
	Byte            int8                                 `json:"byte"`
	Comp            string                               `json:"comp"`
	ConstantKwd     string                               `json:"constant_kwd"`
	Date            AllDateDate                          `json:"date"`
	DateNano        AllDateNanoDate                      `json:"dateNano"`
	DateRange       fielddatatype.Range[builtin.Default] `json:"date_range"`
	DenseVector     []float64                            `json:"dense_vector"`
	Double          float64                              `json:"double"`
	DoubleRange     fielddatatype.Range[float64]         `json:"double_range"`
	Flattened       map[string]any                       `json:"flattened"`
	Float           float32                              `json:"float"`
	FloatRange      fielddatatype.Range[float32]         `json:"float_range"`
	Geopoint        fielddatatype.GeoPoint               `json:"geopoint"`
	Geoshape        fielddatatype.GeoShape               `json:"geoshape"`
	HalfFloat       float32                              `json:"half_float"`
	Histogram       map[string]any                       `json:"histogram"`
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

func (d All) ToRaw() AllRaw {
	return AllRaw{
		Agg:             elastic.FromSingle(d.Agg),
		Blob:            elastic.FromSingle(d.Blob),
		Bool:            elastic.FromSingle(d.Bool),
		Byte:            elastic.FromSingle(d.Byte),
		Comp:            elastic.FromSingle(d.Comp),
		ConstantKwd:     elastic.FromSingle(d.ConstantKwd),
		Date:            elastic.FromSingle(d.Date),
		DateNano:        elastic.FromSingle(d.DateNano),
		DateRange:       elastic.FromSingle(d.DateRange),
		DenseVector:     elastic.FromSingle(d.DenseVector),
		Double:          elastic.FromSingle(d.Double),
		DoubleRange:     elastic.FromSingle(d.DoubleRange),
		Flattened:       elastic.FromSingle(d.Flattened),
		Float:           elastic.FromSingle(d.Float),
		FloatRange:      elastic.FromSingle(d.FloatRange),
		Geopoint:        elastic.FromSingle(d.Geopoint),
		Geoshape:        elastic.FromSingle(d.Geoshape),
		HalfFloat:       elastic.FromSingle(d.HalfFloat),
		Histogram:       elastic.FromSingle(d.Histogram),
		Integer:         elastic.FromSingle(d.Integer),
		IntegerRange:    elastic.FromSingle(d.IntegerRange),
		IpAddr:          elastic.FromSingle(d.IpAddr),
		IpRange:         elastic.FromSingle(d.IpRange),
		Join:            elastic.FromSingle(d.Join),
		Kwd:             elastic.FromSingle(d.Kwd),
		Long:            elastic.FromSingle(d.Long),
		LongRange:       elastic.FromSingle(d.LongRange),
		Nested:          elastic.FromSingle(d.Nested.ToRaw()),
		Object:          elastic.FromSingle(d.Object.ToRaw()),
		Point:           elastic.FromSingle(d.Point),
		Query:           elastic.FromSingle(d.Query),
		RankFeature:     elastic.FromSingle(d.RankFeature),
		RankFeatures:    elastic.FromSingle(d.RankFeatures),
		ScaledFloat:     elastic.FromSingle(d.ScaledFloat),
		SearchAsYouType: elastic.FromSingle(d.SearchAsYouType),
		Shape:           elastic.FromSingle(d.Shape),
		Short:           elastic.FromSingle(d.Short),
		Text:            elastic.FromSingle(d.Text),
		TextWTokenCount: elastic.FromSingle(d.TextWTokenCount),
		UnsignedLong:    elastic.FromSingle(d.UnsignedLong),
		Version:         elastic.FromSingle(d.Version),
		Wildcard:        elastic.FromSingle(d.Wildcard),
	}
}

type AllRaw struct {
	Agg             elastic.Elastic[fielddatatype.AggregateMetricDouble]  `json:"agg"`
	Alias           elastic.Elastic[any]                                  `json:"alias"`
	Blob            elastic.Elastic[[]byte]                               `json:"blob"`
	Bool            elastic.Elastic[fielddatatype.Boolean]                `json:"bool"`
	Byte            elastic.Elastic[int8]                                 `json:"byte"`
	Comp            elastic.Elastic[string]                               `json:"comp"`
	ConstantKwd     elastic.Elastic[string]                               `json:"constant_kwd"`
	Date            elastic.Elastic[AllDateDate]                          `json:"date"`
	DateNano        elastic.Elastic[AllDateNanoDate]                      `json:"dateNano"`
	DateRange       elastic.Elastic[fielddatatype.Range[builtin.Default]] `json:"date_range"`
	DenseVector     elastic.Elastic[[]float64]                            `json:"dense_vector"`
	Double          elastic.Elastic[float64]                              `json:"double"`
	DoubleRange     elastic.Elastic[fielddatatype.Range[float64]]         `json:"double_range"`
	Flattened       elastic.Elastic[map[string]any]                       `json:"flattened"`
	Float           elastic.Elastic[float32]                              `json:"float"`
	FloatRange      elastic.Elastic[fielddatatype.Range[float32]]         `json:"float_range"`
	Geopoint        elastic.Elastic[fielddatatype.GeoPoint]               `json:"geopoint"`
	Geoshape        elastic.Elastic[fielddatatype.GeoShape]               `json:"geoshape"`
	HalfFloat       elastic.Elastic[float32]                              `json:"half_float"`
	Histogram       elastic.Elastic[map[string]any]                       `json:"histogram"`
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

func (d AllRaw) ToPlain() All {
	return All{
		Agg:             d.Agg.ValueSingle(),
		Alias:           d.Alias.ValueSingle(),
		Blob:            d.Blob.ValueSingle(),
		Bool:            d.Bool.ValueSingle(),
		Byte:            d.Byte.ValueSingle(),
		Comp:            d.Comp.ValueSingle(),
		ConstantKwd:     d.ConstantKwd.ValueSingle(),
		Date:            d.Date.ValueSingle(),
		DateNano:        d.DateNano.ValueSingle(),
		DateRange:       d.DateRange.ValueSingle(),
		DenseVector:     d.DenseVector.ValueSingle(),
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
		Nested:          d.Nested.ValueSingle().ToPlain(),
		Object:          d.Object.ValueSingle().ToPlain(),
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
// It also implements json.Marshaler. As JSON representation it will be marshaled into
// 2006-01-02 15:04:05
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
// It also implements json.Marshaler. As JSON representation it will be marshaled into
// 2006-01-02T15:04:05.000000000Z0700
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

func (d AllNestedObject) ToRaw() AllNestedObjectRaw {
	return AllNestedObjectRaw{
		Age:  elastic.FromSingle(d.Age),
		Name: elastic.FromSingle(d.Name.ToRaw()),
	}
}

type AllNestedObjectRaw struct {
	Age  elastic.Elastic[int32]                  `json:"age"`
	Name elastic.Elastic[AllNestedNameObjectRaw] `json:"name"`
}

func (d AllNestedObjectRaw) ToPlain() AllNestedObject {
	return AllNestedObject{
		Age:  d.Age.ValueSingle(),
		Name: d.Name.ValueSingle().ToPlain(),
	}
}

type AllNestedNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func (d AllNestedNameObject) ToRaw() AllNestedNameObjectRaw {
	return AllNestedNameObjectRaw{
		First: elastic.FromSingle(d.First),
		Last:  elastic.FromSingle(d.Last),
	}
}

type AllNestedNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

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

func (d AllObjectObject) ToRaw() AllObjectObjectRaw {
	return AllObjectObjectRaw{
		Age:  elastic.FromSingle(d.Age),
		Name: elastic.FromSingle(d.Name.ToRaw()),
	}
}

type AllObjectObjectRaw struct {
	Age  elastic.Elastic[int32]                  `json:"age"`
	Name elastic.Elastic[AllObjectNameObjectRaw] `json:"name"`
}

func (d AllObjectObjectRaw) ToPlain() AllObjectObject {
	return AllObjectObject{
		Age:  d.Age.ValueSingle(),
		Name: d.Name.ValueSingle().ToPlain(),
	}
}

type AllObjectNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func (d AllObjectNameObject) ToRaw() AllObjectNameObjectRaw {
	return AllObjectNameObjectRaw{
		First: elastic.FromSingle(d.First),
		Last:  elastic.FromSingle(d.Last),
	}
}

type AllObjectNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

func (d AllObjectNameObjectRaw) ToPlain() AllObjectNameObject {
	return AllObjectNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}
