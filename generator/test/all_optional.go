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

type AllOptional struct {
	Agg             *fielddatatype.AggregateMetricDouble    `json:"agg"`
	Alias           *struct{}                               `json:"alias,omitempty"`
	Blob            *[][]byte                               `json:"blob"`
	Bool            *[]fielddatatype.Boolean                `json:"bool"`
	Byte            *[]int8                                 `json:"byte"`
	Comp            *[]string                               `json:"comp,omitempty"`
	ConstantKwd     *[]string                               `json:"constant_kwd,omitempty"`
	Date            *[]AllOptionalDateDate                  `json:"date"`
	DateNano        *[]AllOptionalDateNanoDate              `json:"dateNano"`
	DateRange       *[]fielddatatype.Range[builtin.Default] `json:"date_range"`
	DenseVector     *[3]float64                             `json:"dense_vector,omitempty"`
	Double          *[]float64                              `json:"double"`
	DoubleRange     *[]fielddatatype.Range[float64]         `json:"double_range"`
	Flattened       *[]map[string]any                       `json:"flattened"`
	Float           *[]float32                              `json:"float"`
	FloatRange      *[]fielddatatype.Range[float32]         `json:"float_range"`
	Geopoint        *[]fielddatatype.GeoPoint               `json:"geopoint"`
	Geoshape        *[]fielddatatype.GeoShape               `json:"geoshape"`
	HalfFloat       *[]float32                              `json:"half_float"`
	Histogram       *fielddatatype.Histogram                `json:"histogram"`
	Integer         *[]int32                                `json:"integer"`
	IntegerRange    *[]fielddatatype.Range[int32]           `json:"integer_range"`
	IpAddr          *[]netip.Addr                           `json:"ip_addr"`
	IpRange         *[]fielddatatype.Range[netip.Addr]      `json:"ip_range"`
	Join            *map[string]any                         `json:"join,omitempty"`
	Kwd             *[]string                               `json:"kwd"`
	Long            *[]int64                                `json:"long"`
	LongRange       *[]fielddatatype.Range[int64]           `json:"long_range"`
	Nested          *[]AllOptionalNestedObject              `json:"nested"`
	Object          *[]AllOptionalObjectObject              `json:"object"`
	Point           *[]map[string]any                       `json:"point"`
	Query           *map[string]any                         `json:"query,omitempty"`
	RankFeature     *float64                                `json:"rank_feature"`
	RankFeatures    *[]map[string]float64                   `json:"rank_features,omitempty"`
	ScaledFloat     *[]float64                              `json:"scaled_float"`
	SearchAsYouType *[]string                               `json:"search_as_you_type"`
	Shape           *[]fielddatatype.GeoShape               `json:"shape"`
	Short           *[]int16                                `json:"short"`
	Text            *[]string                               `json:"text"`
	TextWTokenCount *[]string                               `json:"text_w_token_count"`
	UnsignedLong    *[]uint64                               `json:"unsigned_long"`
	Version         *[]string                               `json:"version"`
	Wildcard        *[]string                               `json:"wildcard"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptional) ToRaw() AllOptionalRaw {
	return AllOptionalRaw{
		Agg:             gentypehelper.MapSingleOptionalValueToElastic[fielddatatype.AggregateMetricDouble](d.Agg),
		Blob:            gentypehelper.MapMultipleOptionalValueToElastic[[]byte](d.Blob),
		Bool:            gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.Boolean](d.Bool),
		Byte:            gentypehelper.MapMultipleOptionalValueToElastic[int8](d.Byte),
		Comp:            gentypehelper.MapPlainMultiplePointerToUndefinedElastic[string](d.Comp),
		ConstantKwd:     gentypehelper.MapPlainMultiplePointerToUndefinedElastic[string](d.ConstantKwd),
		Date:            gentypehelper.MapMultipleOptionalValueToElastic[AllOptionalDateDate](d.Date),
		DateNano:        gentypehelper.MapMultipleOptionalValueToElastic[AllOptionalDateNanoDate](d.DateNano),
		DateRange:       gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.Range[builtin.Default]](d.DateRange),
		DenseVector:     sliceund.FromPointer(d.DenseVector),
		Double:          gentypehelper.MapMultipleOptionalValueToElastic[float64](d.Double),
		DoubleRange:     gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.Range[float64]](d.DoubleRange),
		Flattened:       gentypehelper.MapMultipleOptionalValueToElastic[map[string]any](d.Flattened),
		Float:           gentypehelper.MapMultipleOptionalValueToElastic[float32](d.Float),
		FloatRange:      gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.Range[float32]](d.FloatRange),
		Geopoint:        gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.GeoPoint](d.Geopoint),
		Geoshape:        gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.GeoShape](d.Geoshape),
		HalfFloat:       gentypehelper.MapMultipleOptionalValueToElastic[float32](d.HalfFloat),
		Histogram:       gentypehelper.MapSingleOptionalValueToElastic[fielddatatype.Histogram](d.Histogram),
		Integer:         gentypehelper.MapMultipleOptionalValueToElastic[int32](d.Integer),
		IntegerRange:    gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.Range[int32]](d.IntegerRange),
		IpAddr:          gentypehelper.MapMultipleOptionalValueToElastic[netip.Addr](d.IpAddr),
		IpRange:         gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.Range[netip.Addr]](d.IpRange),
		Join:            gentypehelper.MapPlainPointerToUndefinedElastic[map[string]any](d.Join),
		Kwd:             gentypehelper.MapMultipleOptionalValueToElastic[string](d.Kwd),
		Long:            gentypehelper.MapMultipleOptionalValueToElastic[int64](d.Long),
		LongRange:       gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.Range[int64]](d.LongRange),
		Nested:          gentypehelper.MapPlainMultipleOptionalToRawElastic[AllOptionalNestedObjectRaw](d.Nested),
		Object:          gentypehelper.MapPlainMultipleOptionalToRawElastic[AllOptionalObjectObjectRaw](d.Object),
		Point:           gentypehelper.MapMultipleOptionalValueToElastic[map[string]any](d.Point),
		Query:           gentypehelper.MapPlainPointerToUndefinedElastic[map[string]any](d.Query),
		RankFeature:     gentypehelper.MapSingleOptionalValueToElastic[float64](d.RankFeature),
		RankFeatures:    gentypehelper.MapPlainMultiplePointerToUndefinedElastic[map[string]float64](d.RankFeatures),
		ScaledFloat:     gentypehelper.MapMultipleOptionalValueToElastic[float64](d.ScaledFloat),
		SearchAsYouType: gentypehelper.MapMultipleOptionalValueToElastic[string](d.SearchAsYouType),
		Shape:           gentypehelper.MapMultipleOptionalValueToElastic[fielddatatype.GeoShape](d.Shape),
		Short:           gentypehelper.MapMultipleOptionalValueToElastic[int16](d.Short),
		Text:            gentypehelper.MapMultipleOptionalValueToElastic[string](d.Text),
		TextWTokenCount: gentypehelper.MapMultipleOptionalValueToElastic[string](d.TextWTokenCount),
		UnsignedLong:    gentypehelper.MapMultipleOptionalValueToElastic[uint64](d.UnsignedLong),
		Version:         gentypehelper.MapMultipleOptionalValueToElastic[string](d.Version),
		Wildcard:        gentypehelper.MapMultipleOptionalValueToElastic[string](d.Wildcard),
	}
}

type AllOptionalRaw struct {
	Agg             elastic.Elastic[fielddatatype.AggregateMetricDouble]  `json:"agg,omitempty"`
	Alias           elastic.Elastic[*struct{}]                            `json:"alias,omitempty"`
	Blob            elastic.Elastic[[]byte]                               `json:"blob,omitempty"`
	Bool            elastic.Elastic[fielddatatype.Boolean]                `json:"bool,omitempty"`
	Byte            elastic.Elastic[int8]                                 `json:"byte,omitempty"`
	Comp            elastic.Elastic[string]                               `json:"comp,omitempty"`
	ConstantKwd     elastic.Elastic[string]                               `json:"constant_kwd,omitempty"`
	Date            elastic.Elastic[AllOptionalDateDate]                  `json:"date,omitempty"`
	DateNano        elastic.Elastic[AllOptionalDateNanoDate]              `json:"dateNano,omitempty"`
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
	Nested          elastic.Elastic[AllOptionalNestedObjectRaw]           `json:"nested,omitempty"`
	Object          elastic.Elastic[AllOptionalObjectObjectRaw]           `json:"object,omitempty"`
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
func (d AllOptionalRaw) ToPlain() AllOptional {
	return AllOptional{
		Agg:             d.Agg.Pointer(),
		Blob:            gentypehelper.MapElasticToMultipleValueOptional[[]byte](d.Blob),
		Bool:            gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.Boolean](d.Bool),
		Byte:            gentypehelper.MapElasticToMultipleValueOptional[int8](d.Byte),
		Comp:            gentypehelper.MapElasticToMultipleValueOptional[string](d.Comp),
		ConstantKwd:     gentypehelper.MapElasticToMultipleValueOptional[string](d.ConstantKwd),
		Date:            gentypehelper.MapElasticToMultipleValueOptional[AllOptionalDateDate](d.Date),
		DateNano:        gentypehelper.MapElasticToMultipleValueOptional[AllOptionalDateNanoDate](d.DateNano),
		DateRange:       gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.Range[builtin.Default]](d.DateRange),
		DenseVector:     d.DenseVector.Pointer(),
		Double:          gentypehelper.MapElasticToMultipleValueOptional[float64](d.Double),
		DoubleRange:     gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.Range[float64]](d.DoubleRange),
		Flattened:       gentypehelper.MapElasticToMultipleValueOptional[map[string]any](d.Flattened),
		Float:           gentypehelper.MapElasticToMultipleValueOptional[float32](d.Float),
		FloatRange:      gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.Range[float32]](d.FloatRange),
		Geopoint:        gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.GeoPoint](d.Geopoint),
		Geoshape:        gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.GeoShape](d.Geoshape),
		HalfFloat:       gentypehelper.MapElasticToMultipleValueOptional[float32](d.HalfFloat),
		Histogram:       d.Histogram.Pointer(),
		Integer:         gentypehelper.MapElasticToMultipleValueOptional[int32](d.Integer),
		IntegerRange:    gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.Range[int32]](d.IntegerRange),
		IpAddr:          gentypehelper.MapElasticToMultipleValueOptional[netip.Addr](d.IpAddr),
		IpRange:         gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.Range[netip.Addr]](d.IpRange),
		Join:            d.Join.Pointer(),
		Kwd:             gentypehelper.MapElasticToMultipleValueOptional[string](d.Kwd),
		Long:            gentypehelper.MapElasticToMultipleValueOptional[int64](d.Long),
		LongRange:       gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.Range[int64]](d.LongRange),
		Nested:          gentypehelper.MapElasticToPlainMultipleOptional[AllOptionalNestedObject](d.Nested),
		Object:          gentypehelper.MapElasticToPlainMultipleOptional[AllOptionalObjectObject](d.Object),
		Point:           gentypehelper.MapElasticToMultipleValueOptional[map[string]any](d.Point),
		Query:           d.Query.Pointer(),
		RankFeature:     d.RankFeature.Pointer(),
		RankFeatures:    gentypehelper.MapElasticToMultipleValueOptional[map[string]float64](d.RankFeatures),
		ScaledFloat:     gentypehelper.MapElasticToMultipleValueOptional[float64](d.ScaledFloat),
		SearchAsYouType: gentypehelper.MapElasticToMultipleValueOptional[string](d.SearchAsYouType),
		Shape:           gentypehelper.MapElasticToMultipleValueOptional[fielddatatype.GeoShape](d.Shape),
		Short:           gentypehelper.MapElasticToMultipleValueOptional[int16](d.Short),
		Text:            gentypehelper.MapElasticToMultipleValueOptional[string](d.Text),
		TextWTokenCount: gentypehelper.MapElasticToMultipleValueOptional[string](d.TextWTokenCount),
		UnsignedLong:    gentypehelper.MapElasticToMultipleValueOptional[uint64](d.UnsignedLong),
		Version:         gentypehelper.MapElasticToMultipleValueOptional[string](d.Version),
		Wildcard:        gentypehelper.MapElasticToMultipleValueOptional[string](d.Wildcard),
	}
}

// AllOptionalDateDate represents the date or the date_nanos mapping field type.
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
type AllOptionalDateDate time.Time

var parserAllOptionalDateDate = estime.FromGoTimeLayoutUnsafe(
	[]string{
		"2006-01-02 15:04:05",
		"2006-01-02",
	},
	"epoch_millis",
)

// String implements fmt.Stringer
func (t AllOptionalDateDate) String() string {
	return parserAllOptionalDateDate.FormatString(time.Time(t), 0)
}

// MarshalJSON implements json.Marshaler
func (t AllOptionalDateDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (t *AllOptionalDateDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	tt, err := parserAllOptionalDateDate.ParseJson(data)
	if err != nil {
		return err
	}
	*t = AllOptionalDateDate(tt)
	return nil
}

// AllOptionalDateNanoDate represents the date or the date_nanos mapping field type.
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
type AllOptionalDateNanoDate time.Time

var parserAllOptionalDateNanoDate = estime.FromGoTimeLayoutUnsafe(
	[]string{
		"2006-01-02T15:04:05.000000000Z0700",
		"2006-01-02",
	},
	"epoch_second",
)

// String implements fmt.Stringer
func (t AllOptionalDateNanoDate) String() string {
	return parserAllOptionalDateNanoDate.FormatString(time.Time(t), 0)
}

// MarshalJSON implements json.Marshaler
func (t AllOptionalDateNanoDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (t *AllOptionalDateNanoDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	tt, err := parserAllOptionalDateNanoDate.ParseJson(data)
	if err != nil {
		return err
	}
	*t = AllOptionalDateNanoDate(tt)
	return nil
}

type AllOptionalNestedObject struct {
	Age  *[]int32                       `json:"age"`
	Name *[]AllOptionalNestedNameObject `json:"name"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptionalNestedObject) ToRaw() AllOptionalNestedObjectRaw {
	return AllOptionalNestedObjectRaw{
		Age:  gentypehelper.MapMultipleOptionalValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainMultipleOptionalToRawElastic[AllOptionalNestedNameObjectRaw](d.Name),
	}
}

type AllOptionalNestedObjectRaw struct {
	Age  elastic.Elastic[int32]                          `json:"age,omitempty"`
	Name elastic.Elastic[AllOptionalNestedNameObjectRaw] `json:"name,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptionalNestedObjectRaw) ToPlain() AllOptionalNestedObject {
	return AllOptionalNestedObject{
		Age:  gentypehelper.MapElasticToMultipleValueOptional[int32](d.Age),
		Name: gentypehelper.MapElasticToPlainMultipleOptional[AllOptionalNestedNameObject](d.Name),
	}
}

type AllOptionalNestedNameObject struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptionalNestedNameObject) ToRaw() AllOptionalNestedNameObjectRaw {
	return AllOptionalNestedNameObjectRaw{
		First: gentypehelper.MapMultipleOptionalValueToElastic[string](d.First),
		Last:  gentypehelper.MapMultipleOptionalValueToElastic[string](d.Last),
	}
}

type AllOptionalNestedNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first,omitempty"`
	Last  elastic.Elastic[string] `json:"last,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptionalNestedNameObjectRaw) ToPlain() AllOptionalNestedNameObject {
	return AllOptionalNestedNameObject{
		First: gentypehelper.MapElasticToMultipleValueOptional[string](d.First),
		Last:  gentypehelper.MapElasticToMultipleValueOptional[string](d.Last),
	}
}

type AllOptionalObjectObject struct {
	Age  *[]int32                       `json:"age"`
	Name *[]AllOptionalObjectNameObject `json:"name"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptionalObjectObject) ToRaw() AllOptionalObjectObjectRaw {
	return AllOptionalObjectObjectRaw{
		Age:  gentypehelper.MapMultipleOptionalValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainMultipleOptionalToRawElastic[AllOptionalObjectNameObjectRaw](d.Name),
	}
}

type AllOptionalObjectObjectRaw struct {
	Age  elastic.Elastic[int32]                          `json:"age,omitempty"`
	Name elastic.Elastic[AllOptionalObjectNameObjectRaw] `json:"name,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptionalObjectObjectRaw) ToPlain() AllOptionalObjectObject {
	return AllOptionalObjectObject{
		Age:  gentypehelper.MapElasticToMultipleValueOptional[int32](d.Age),
		Name: gentypehelper.MapElasticToPlainMultipleOptional[AllOptionalObjectNameObject](d.Name),
	}
}

type AllOptionalObjectNameObject struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptionalObjectNameObject) ToRaw() AllOptionalObjectNameObjectRaw {
	return AllOptionalObjectNameObjectRaw{
		First: gentypehelper.MapMultipleOptionalValueToElastic[string](d.First),
		Last:  gentypehelper.MapMultipleOptionalValueToElastic[string](d.Last),
	}
}

type AllOptionalObjectNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first,omitempty"`
	Last  elastic.Elastic[string] `json:"last,omitempty"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AllOptionalObjectNameObjectRaw) ToPlain() AllOptionalObjectNameObject {
	return AllOptionalObjectNameObject{
		First: gentypehelper.MapElasticToMultipleValueOptional[string](d.First),
		Last:  gentypehelper.MapElasticToMultipleValueOptional[string](d.Last),
	}
}
