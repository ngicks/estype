package generator

import (
	"strings"

	"github.com/ngicks/estype/spec/mapping"
	"github.com/ngicks/und/option"
)

type GenerateTypeName func(fieldNames []string, typeName mapping.EsType) string

var (
	ChainFieldName GenerateTypeName = func(fieldNames []string, typeName mapping.EsType) string {
		names := make([]string, len(fieldNames)+1)
		for idx, v := range fieldNames {
			names[idx] = pascalCase(v)
		}
		names[len(names)-1] = pascalCase(string(typeName))
		return strings.Join(names, "")
	}
)

type MappingOption map[string]PropertyOption

type PropertyOption struct {
	TypeName                  string
	IsOptional                option.Option[bool]
	IsSingle                  option.Option[bool]
	PreferStringBoolean       option.Option[bool]
	PreferMarshalDateToNumber option.Option[bool]
	Child                     MappingOption
}

func (o PropertyOption) GetTypeName(
	generateTypeName func() string,
) string {
	if o.TypeName != "" {
		return o.TypeName
	}
	return generateTypeName()
}

type DefaultOption struct {
	IsOptional                option.Option[bool]
	IsSingle                  option.Option[bool]
	PreferStringBoolean       option.Option[bool]
	PreferMarshalDateToNumber option.Option[bool]
	PerTypDefault             map[mapping.EsType]TypeOption
}

type TypeOption struct {
	IsOptional option.Option[bool]
	IsSingle   option.Option[bool]
}

var (
	trueOp = option.Some[bool](true)
)

func GetDefaultTypeOption() map[mapping.EsType]TypeOption {
	return map[mapping.EsType]TypeOption{
		mapping.AggregateMetricDouble: {},
		mapping.FieldAlias: {
			IsSingle: trueOp,
		},
		mapping.Binary: {},
		mapping.Boolean: {
			IsSingle: trueOp,
		},
		mapping.Completion: {},
		mapping.Date: {
			IsSingle: trueOp,
		},
		mapping.DateNanos: {
			IsSingle: trueOp,
		},
		mapping.DenseVector: {},
		mapping.Flattened:   {},
		mapping.GeoPoint:    {},
		mapping.GeoShape:    {},
		mapping.Histogram:   {},
		mapping.Ip:          {},
		mapping.Join:        {},
		mapping.Nested:      {},
		mapping.Object: {
			IsSingle: trueOp,
		},
		mapping.Percolator:      {},
		mapping.Point:           {},
		mapping.RankFeature:     {},
		mapping.RankFeatures:    {},
		mapping.SearchAsYouType: {},
		mapping.Shape:           {},
		mapping.TokenCount: {
			IsSingle: trueOp,
		},
		mapping.Version: {
			IsSingle: trueOp,
		},
		mapping.Keyword: {},
		mapping.ConstantKeyword: {
			IsSingle: trueOp,
		},
		mapping.Wildcard:           {},
		mapping.Text:               {},
		mapping.LongNumber:         {},
		mapping.IntegerNumber:      {},
		mapping.ShortNumber:        {},
		mapping.ByteNumber:         {},
		mapping.DoubleNumber:       {},
		mapping.FloatNumber:        {},
		mapping.HalfFloatNumber:    {},
		mapping.ScaledFloatNumber:  {},
		mapping.UnsignedLongNumber: {},
		mapping.IntegerRange:       {},
		mapping.FloatRange:         {},
		mapping.LongRange:          {},
		mapping.DoubleRange:        {},
		mapping.DateRange:          {},
		mapping.IpRange:            {},
	}
}
