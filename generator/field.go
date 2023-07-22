package generator

import (
	"github.com/ngicks/estype/spec/mapping"
)

const (
	fielddatatypeQual = "github.com/ngicks/estype/fielddatatype"
)

const (
	anyMap     = "map[string]any"
	float64Map = "map[string]float64"
)

// genField generates a type for input property.
// Input prop must be one that can not be nested (other than Object or Nested types).
func genField(ctx *GeneratorContext, dryRun bool) (typeName TypeId) {
	if ty, ok := fieldTypeTable[mapping.GetTypeName(ctx.localState.prop)]; ok {
		return ty
	}

	switch x := ctx.localState.prop.Val.(type) {
	case mapping.AggregateMetricDoubleProperty:
		gen := genAggregateMetricDouble(x)
		return gen
	case mapping.BooleanProperty:
		if ctx.PreferStringBoolean() {
			return TypeId{
				Qualifier: fielddatatypeQual,
				Id:        "BooleanStr",
			}
		} else {
			return TypeId{
				Qualifier: fielddatatypeQual,
				Id:        "Boolean",
			}
		}
	case mapping.DateProperty, mapping.DateNanosProperty:
		return genDate(ctx, dryRun)
	case mapping.DenseVectorProperty:
		return genDenseVector(x)
	case mapping.DateRangeProperty, mapping.DoubleRangeProperty, mapping.FloatRangeProperty,
		mapping.IntegerRangeProperty, mapping.IpRangeProperty, mapping.LongRangeProperty:
		return genRange(ctx, dryRun)
	}

	// return any for unknown types.
	return TypeId{Id: "any"}
}

var fieldTypeTable = map[mapping.EsType]TypeId{
	mapping.FieldAlias:      {NonWritable: true},
	mapping.Binary:          {Id: "[]byte"},
	mapping.Completion:      {Id: "string"},
	mapping.Flattened:       {Id: anyMap},
	mapping.GeoPoint:        {Id: "GeoPoint", Qualifier: fielddatatypeQual},
	mapping.GeoShape:        {Id: "GeoShape", Qualifier: fielddatatypeQual},
	mapping.Ip:              {Id: "Addr", Qualifier: "net/netip"},
	mapping.Histogram:       {Id: "Histogram", Qualifier: fielddatatypeQual},
	mapping.Join:            {Id: anyMap}, // TODO: implement
	mapping.Percolator:      {Id: anyMap}, // TODO: implement
	mapping.Point:           {Id: anyMap}, // TODO: implement
	mapping.RankFeature:     {Id: "float64"},
	mapping.RankFeatures:    {Id: float64Map},
	mapping.SearchAsYouType: {Id: "string"},
	mapping.Shape:           {Id: "GeoShape", Qualifier: fielddatatypeQual},
	mapping.TokenCount:      {Id: "int64"},
	mapping.Version:         {Id: "string"}, // should this be sem ver package?
	mapping.Keyword:         {Id: "string"},
	// The field can be stored if and only if value is same as specified in param.
	// Should this field also be considered non writable?
	mapping.ConstantKeyword: {Id: "string"},
	mapping.Wildcard:        {Id: "string"},
	mapping.Text:            {Id: "string"},
	mapping.MatchOnlyText:   {Id: "string"},
	// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/number.html
	mapping.LongNumber:         {Id: "int64"},
	mapping.IntegerNumber:      {Id: "int32"},
	mapping.ShortNumber:        {Id: "int16"},
	mapping.ByteNumber:         {Id: "int8"}, // The doc says it ranges -128 to 127. It's not the go built-in byte. Rather, it is a typical char type.
	mapping.DoubleNumber:       {Id: "float64"},
	mapping.FloatNumber:        {Id: "float32"},
	mapping.HalfFloatNumber:    {Id: "float32"}, // TODO: use float16 package?
	mapping.UnsignedLongNumber: {Id: "uint64"},
	mapping.ScaledFloatNumber:  {Id: "float64"},
}
