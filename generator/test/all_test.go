package test

import (
	"encoding/json"
	"net/netip"
	"testing"
	"time"

	"github.com/go-spatial/geom"
	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/estype/fielddatatype"
	"github.com/ngicks/estype/fielddatatype/estime/builtin"
	elastic "github.com/ngicks/und/elastic"
	undefinedable "github.com/ngicks/und/undefinedable"
	"github.com/stretchr/testify/require"
)

func escape[T any](v T) *T {
	return &v
}

func backConvert[T interface{ ToRaw() U }, U interface{ ToPlain() T }](v any) any {
	return v.(T).ToRaw().ToPlain()
}

func toPlain[T any, U interface{ ToPlain() T }](v any) any {
	return v.(U).ToPlain()
}

type losslessConversionTestCase struct {
	inputPlain         any
	unmarshalTargetRaw any
	backConvert        func(any) any
	toPlain            func(any) any
}

func TestLosslessConversion(t *testing.T) {
	require := require.New(t)

	for _, tc := range []losslessConversionTestCase{
		{
			inputPlain:         sampleAll,
			unmarshalTargetRaw: &AllRaw{},
			backConvert:        backConvert[All, AllRaw],
			toPlain:            toPlain[All, *AllRaw],
		},
		{
			inputPlain:         sampleAllOptional,
			unmarshalTargetRaw: &AllOptionalRaw{},
			backConvert:        backConvert[AllOptional, AllOptionalRaw],
			toPlain:            toPlain[AllOptional, *AllOptionalRaw],
		},
		{
			inputPlain:         sampleConversion,
			unmarshalTargetRaw: &ConversionRaw{},
			backConvert:        backConvert[Conversion, ConversionRaw],
			toPlain:            toPlain[Conversion, *ConversionRaw],
		},
		{
			inputPlain:         sampleDynamic,
			unmarshalTargetRaw: &DynamicRaw{},
			backConvert:        backConvert[Dynamic, DynamicRaw],
			toPlain:            toPlain[Dynamic, *DynamicRaw],
		},
		{
			inputPlain:         sampleAddtionalPropEscape,
			unmarshalTargetRaw: &AddtionalPropEscapeRaw{},
			backConvert:        backConvert[AddtionalPropEscape, AddtionalPropEscapeRaw],
			toPlain:            toPlain[AddtionalPropEscape, *AddtionalPropEscapeRaw],
		},
	} {
		marshaled1st, err := json.Marshal(tc.inputPlain)
		require.NoError(err)
		marshaled2nd, err := json.Marshal(tc.backConvert(tc.inputPlain))
		require.NoError(err)

		err = json.Unmarshal(marshaled1st, tc.unmarshalTargetRaw)
		require.NoError(err)
		marshaled3rd, err := json.Marshal(tc.toPlain(tc.unmarshalTargetRaw))
		require.NoError(err)

		t.Logf("%s\n", marshaled1st)
		t.Logf("%s\n", marshaled2nd)
		t.Logf("%s\n", marshaled3rd)
		t.Logf("\n")
		diff := cmp.Diff(string(marshaled1st), string(marshaled2nd))
		if diff != "" {
			t.Fatalf("not equal. diff = %s", diff)
		}
		diff = cmp.Diff(string(marshaled2nd), string(marshaled3rd))
		if diff != "" {
			t.Fatalf("not equal. diff = %s", diff)
		}
	}
}

var sampleTime = time.Date(2023, 10, 26, 2, 34, 52, 123456789, time.UTC)

var sampleAll = All{
	Agg: fielddatatype.AggregateMetricDouble{
		Min:        123,
		Max:        1270853,
		Sum:        503,
		ValueCount: 2178,
	},
	Blob:        []byte(`foobarbaz`),
	Bool:        fielddatatype.Boolean(true),
	Byte:        int8(12),
	Comp:        "jdwioujaocksjml.nhmnl.c",
	ConstantKwd: "debug",
	Date:        AllDateDate(sampleTime),
	DateNano:    AllDateNanoDate(sampleTime),
	DateRange: fielddatatype.Range[builtin.Default]{
		Gte: escape(builtin.Default(time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC))),
		Lt:  escape(builtin.Default(time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC))),
	},
	DenseVector: [3]float64{16, 15, 14},
	Double:      float64(68),
	DoubleRange: fielddatatype.Range[float64]{
		Gte: escape(10.1),
		Lt:  escape(20.1),
	},
	Flattened: map[string]interface{}{
		"priority": "urgent",
		"release":  []any{"v1.2.5", "v1.3.0"},
		"timestamp": map[string]any{
			"created": float64(1541458026),
			"closed":  float64(1541457010),
		},
	},
	Float: float32(357.3209),
	FloatRange: fielddatatype.Range[float32]{
		Gte: escape[float32](10.1),
		Lt:  escape[float32](20.1),
	},
	Geopoint: fielddatatype.GeoPoint{
		Lat: 41.12,
		Lon: -71.34,
	},
	Geoshape: fielddatatype.GeoShape{
		Geometry: geom.Point{-77.03653, 38.897676},
	},
	HalfFloat: float32(2131.57),
	Histogram: fielddatatype.Histogram{
		Values: []float64{0.1, 0.2, 0.3, 0.4, 0.5},
		Counts: []int32{3, 7, 23, 12, 6},
	},
	Integer: int32(60),
	IntegerRange: fielddatatype.Range[int32]{
		Gte: escape[int32](10),
		Lt:  escape[int32](20),
	},
	IpAddr: netip.MustParseAddr("192.168.0.1"),
	IpRange: fielddatatype.Range[netip.Addr]{
		Gte: escape(netip.AddrFrom4([4]byte{192, 168, 0, 2})),
		Lt:  escape(netip.AddrFrom4([4]byte{192, 168, 0, 240})),
	},
	Join: (map[string]interface{}{
		"name": "question",
	}),
	Kwd:  "naaaaaaaaaaaaaah",
	Long: int64(210389467827),
	LongRange: fielddatatype.Range[int64]{
		Gte: escape[int64](10),
		Lt:  escape[int64](20),
	},
	Nested: AllNestedObject{
		Age: int32(123),
		Name: AllNestedNameObject{
			First: "john",
			Last:  "doe",
		},
	},
	Object: AllObjectObject{
		Age: int32(123),
		Name: AllObjectNameObject{
			First: "john",
			Last:  "doe",
		},
	},
	Point: (map[string]interface{}{
		"type":        "Point",
		"coordinates": []any{-71.34, 41.12},
	}),
	Query: (map[string]interface{}{
		"match": map[string]any{
			"kwd": "value",
		},
	}),
	RankFeature: float64(124.6),
	RankFeatures: (map[string]float64{
		"politics":  float64(20),
		"economics": 50.8,
	}),
	ScaledFloat:     float64(12315.4798),
	SearchAsYouType: "quick brown fox jump lazy dog",
	Shape: fielddatatype.GeoShape{
		Geometry: geom.Point{-77.03653, 38.897676},
	},
	Short:           int16(2109),
	Text:            "fox fox fox",
	TextWTokenCount: "1208956i;lzcxjo",
	UnsignedLong:    uint64(2109381027538706718),
	Version:         "1.2.7",
	Wildcard:        "8lnmkvlouiejhr02983",
}

var sampleEmptyAllRaw = AllRaw{
	Agg:             elastic.Null[fielddatatype.AggregateMetricDouble](),
	Blob:            elastic.Null[[]byte](),
	Bool:            elastic.Null[fielddatatype.Boolean](),
	Byte:            elastic.Null[int8](),
	Comp:            elastic.Undefined[string](),
	ConstantKwd:     elastic.Undefined[string](),
	Date:            elastic.Null[AllDateDate](),
	DateNano:        elastic.Null[AllDateNanoDate](),
	DateRange:       elastic.Null[fielddatatype.Range[builtin.Default]](),
	DenseVector:     undefinedable.Undefined[[3]float64](),
	Double:          elastic.Null[float64](),
	DoubleRange:     elastic.Null[fielddatatype.Range[float64]](),
	Flattened:       elastic.Null[map[string]any](),
	Float:           elastic.Null[float32](),
	FloatRange:      elastic.Null[fielddatatype.Range[float32]](),
	Geopoint:        elastic.Null[fielddatatype.GeoPoint](),
	Geoshape:        elastic.Null[fielddatatype.GeoShape](),
	HalfFloat:       elastic.Null[float32](),
	Histogram:       elastic.Null[fielddatatype.Histogram](),
	Integer:         elastic.Null[int32](),
	IntegerRange:    elastic.Null[fielddatatype.Range[int32]](),
	IpAddr:          elastic.Null[netip.Addr](),
	IpRange:         elastic.Null[fielddatatype.Range[netip.Addr]](),
	Join:            elastic.Undefined[map[string]any](),
	Kwd:             elastic.Null[string](),
	Long:            elastic.Null[int64](),
	LongRange:       elastic.Null[fielddatatype.Range[int64]](),
	Nested:          elastic.Null[AllNestedObjectRaw](),
	Object:          elastic.Null[AllObjectObjectRaw](),
	Point:           elastic.Null[map[string]any](),
	Query:           elastic.Undefined[map[string]any](),
	RankFeature:     elastic.Null[float64](),
	RankFeatures:    elastic.Undefined[map[string]float64](),
	ScaledFloat:     elastic.Null[float64](),
	SearchAsYouType: elastic.Null[string](),
	Shape:           elastic.Null[fielddatatype.GeoShape](),
	Short:           elastic.Null[int16](),
	Text:            elastic.Null[string](),
	TextWTokenCount: elastic.Null[string](),
	UnsignedLong:    elastic.Null[uint64](),
	Version:         elastic.Null[string](),
	Wildcard:        elastic.Null[string](),
}
