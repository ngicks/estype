package generator

import (
	"fmt"

	"github.com/ngicks/estype/spec/mapping"
)

func genDenseVector(prop mapping.DenseVectorProperty) typeId {
	return typeId{
		Id:            fmt.Sprintf("[%d]float64", prop.Dims),
		AlwaysSingle:  true,
		DisallowArray: true,
		DisallowNull:  true,
	}
}
