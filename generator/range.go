package generator

import (
	"fmt"

	"github.com/ngicks/estype/spec/mapping"
)

func genRange(ctx *generatorContext, dryRun bool) typeId {
	var param typeId
	switch ctx.localState.prop.Val.(type) {
	case mapping.DateRangeProperty:
		param = genDate(ctx, dryRun)
	case mapping.DoubleRangeProperty:
		param = typeId{Id: "float64"}
	case mapping.FloatRangeProperty:
		param = typeId{Id: "float32"}
	case mapping.IntegerRangeProperty:
		param = typeId{Id: "int32"}
	case mapping.IpRangeProperty:
		param = typeId{Id: "Addr", Qualifier: "net/netip"}
	case mapping.LongRangeProperty:
		param = typeId{Id: "int64"}
	default:
		panic(
			fmt.Errorf(
				"unknown type. Range must not be called with context set to one of Range type. input = %T",
				ctx.localState.prop.Val,
			),
		)
	}

	return typeId{
		Id:        "Range",
		Qualifier: fielddatatypeQual,
		TypeParam: []typeId{param},
	}
}
