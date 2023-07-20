package generator

import (
	"fmt"

	"github.com/ngicks/estype/spec/mapping"
)

func Range(ctx *GeneratorContext, dryRun bool) TypeId {
	var param TypeId
	switch ctx.localState.prop.Val.(type) {
	case mapping.DateRangeProperty:
		param = Date(ctx, dryRun)
	case mapping.DoubleRangeProperty:
		param = TypeId{Id: "float64"}
	case mapping.FloatRangeProperty:
		param = TypeId{Id: "float32"}
	case mapping.IntegerRangeProperty:
		param = TypeId{Id: "int32"}
	case mapping.IpRangeProperty:
		param = TypeId{Id: "Addr", Qualifier: "net/netip"}
	case mapping.LongRangeProperty:
		param = TypeId{Id: "int64"}
	default:
		panic(
			fmt.Errorf(
				"unknown type. Range must not be called with context set to one of Range type. input = %T",
				ctx.localState.prop.Val,
			),
		)
	}

	return TypeId{
		Id:        "Range",
		Qualifier: fielddatatypeQual,
		TypeParam: []TypeId{param},
	}
}
