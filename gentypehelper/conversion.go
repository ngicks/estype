package gentypehelper

import "github.com/ngicks/und/sliceund/elastic"

const (
	IdMapPlainPointerToUndefinedElastic         = "MapPlainPointerToUndefinedElastic"
	IdMapPlainMultiplePointerToUndefinedElastic = "MapPlainMultiplePointerToUndefinedElastic"
	IdMapSingleValueToElastic                   = "MapSingleValueToElastic"
	IdMapSingleOptionalValueToElastic           = "MapSingleOptionalValueToElastic"
	IdMapMultipleValueToElastic                 = "MapMultipleValueToElastic"
	IdMapMultipleOptionalValueToElastic         = "MapMultipleOptionalValueToElastic"
	IdMapPlainToRawElastic                      = "MapPlainToRawElastic"
	IdMapPlainOptionalToRawElastic              = "MapPlainOptionalToRawElastic"
	IdMapPlainMultipleToRawElastic              = "MapPlainMultipleToRawElastic"
	IdMapPlainMultipleOptionalToRawElastic      = "MapPlainMultipleOptionalToRawElastic"
	IdMapElasticToPlainSingle                   = "MapElasticToPlainSingle"
	IdMapElasticToPlainSingleOptional           = "MapElasticToPlainSingleOptional"
	IdMapElasticToPlainMultiple                 = "MapElasticToPlainMultiple"
	IdMapElasticToPlainMultipleOptional         = "MapElasticToPlainMultipleOptional"
	IdMapElasticToMultipleValueOptional         = "MapElasticToMultipleValueOptional"
)

func MapPlainPointerToUndefinedElastic[T any](v *T) elastic.Elastic[T] {
	if v == nil {
		return elastic.Undefined[T]()
	}
	return elastic.FromValue(*v)
}

func MapPlainMultiplePointerToUndefinedElastic[T any](v *[]T) elastic.Elastic[T] {
	if v == nil {
		return elastic.Undefined[T]()
	}
	return elastic.FromValues(*v)
}

func MapSingleValueToElastic[T any](v T) elastic.Elastic[T] {
	return elastic.FromValue[T](v)
}

func MapSingleOptionalValueToElastic[T any](v *T) elastic.Elastic[T] {
	return elastic.FromPointer[T](v)
}

func MapMultipleValueToElastic[T any](v []T) elastic.Elastic[T] {
	return elastic.FromValues[T](v)
}

func MapMultipleOptionalValueToElastic[T any](v *[]T) elastic.Elastic[T] {
	if v == nil {
		return elastic.FromValues[T]([]T{})
	}
	return elastic.FromValues[T](*v)
}

func MapPlainToRawElastic[T any, U interface{ ToRaw() T }](v U) elastic.Elastic[T] {
	return elastic.FromValue[T](v.ToRaw())
}

func MapPlainOptionalToRawElastic[T any, U interface{ ToRaw() T }](v *U) elastic.Elastic[T] {
	if v == nil {
		return elastic.Null[T]()
	}
	return elastic.FromValue[T]((*v).ToRaw())
}

func MapPlainMultipleToRawElastic[T any, U interface{ ToRaw() T }](v []U) elastic.Elastic[T] {
	out := make([]T, len(v))
	for idx, vv := range v {
		out[idx] = vv.ToRaw()
	}
	return elastic.FromValues[T](out)
}

func MapPlainMultipleOptionalToRawElastic[T any, U interface{ ToRaw() T }](v *[]U) elastic.Elastic[T] {
	if v == nil {
		return elastic.Null[T]()
	}
	return MapPlainMultipleToRawElastic[T, U](*v)
}

func MapElasticToPlainSingle[T any, U interface{ ToPlain() T }](v elastic.Elastic[U]) T {
	return v.Value().ToPlain()
}

func MapElasticToPlainSingleOptional[T any, U interface{ ToPlain() T }](v elastic.Elastic[U]) *T {
	if v.IsUndefined() || v.IsNull() {
		return nil
	}
	p := MapElasticToPlainSingle[T, U](v)
	return &p
}

func MapElasticToPlainMultiple[T any, U interface{ ToPlain() T }](v elastic.Elastic[U]) []T {
	values := v.Values()
	out := make([]T, len(values))
	for idx, vv := range values {
		out[idx] = vv.ToPlain()
	}
	return out
}

func MapElasticToPlainMultipleOptional[T any, U interface{ ToPlain() T }](v elastic.Elastic[U]) *[]T {
	if v.IsUndefined() || v.IsNull() {
		return nil
	}
	p := MapElasticToPlainMultiple[T, U](v)
	return &p
}

func MapElasticToMultipleValueOptional[T any](v elastic.Elastic[T]) *[]T {
	if v.IsUndefined() || v.IsNull() {
		return nil
	}
	p := v.Values()
	return &p
}
