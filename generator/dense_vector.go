package generator

import (
	"fmt"

	"github.com/ngicks/estype/spec/mapping"
)

func genDenseVector(prop mapping.DenseVectorProperty) TypeId {
	return TypeId{
		Id:            fmt.Sprintf("[%d]float64", prop.Dims),
		AlwaysSingle:  true,
		DisallowArray: true,
		DisallowNull:  true,
	}
}
