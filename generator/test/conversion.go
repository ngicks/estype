package test

import elastic "github.com/ngicks/und/elastic"

func escapeSlice[T any](sl []T) *[]T {
	if sl == nil {
		return nil
	}
	return &sl
}

func mapToPlain[T any, U interface{ ToPlain() T }](l []U) []T {
	if l == nil {
		return nil
	}
	out := make([]T, len(l))
	for i, v := range l {
		out[i] = v.ToPlain()
	}
	return out
}

func mapToRawPointer[T any, U interface{ ToRaw() T }](l *[]U) []T {
	if l == nil {
		return nil
	}
	out := make([]T, len(*l))
	for i, v := range *l {
		out[i] = v.ToRaw()
	}
	return out
}

func mapToRaw[T any, U interface{ ToRaw() T }](l []U) []T {
	if l == nil {
		return nil
	}
	out := make([]T, len(l))
	for i, v := range l {
		out[i] = v.ToRaw()
	}
	return out
}

func escapeValue[T any](v T) *T {
	return &v
}

type Conversion struct {
	MultipleOptional *[]ConversionMultipleOptionalNested `json:"multiple_optional"`
	MultipleRequired []ConversionMultipleRequiredNested  `json:"multiple_required"`
	SingleOptional   *ConversionSingleOptionalObject     `json:"single_optional"`
	SingleRequired   ConversionSingleRequiredObject      `json:"single_required"`
}

func (d Conversion) ToRaw() ConversionRaw {
	return ConversionRaw{
		MultipleOptional: elastic.FromMultiplePointer(escapeSlice(mapToRawPointer[ConversionMultipleOptionalNestedRaw](d.MultipleOptional))),
		MultipleRequired: elastic.FromMultiple(mapToRaw[ConversionMultipleRequiredNestedRaw](d.MultipleRequired)),
		SingleOptional:   elastic.FromSinglePointer(escapeValue(d.SingleOptional.ToRaw())),
		SingleRequired:   elastic.FromSingle(d.SingleRequired.ToRaw()),
	}
}

type ConversionRaw struct {
	MultipleOptional elastic.Elastic[ConversionMultipleOptionalNestedRaw] `json:"multiple_optional"`
	MultipleRequired elastic.Elastic[ConversionMultipleRequiredNestedRaw] `json:"multiple_required"`
	SingleOptional   elastic.Elastic[ConversionSingleOptionalObjectRaw]   `json:"single_optional"`
	SingleRequired   elastic.Elastic[ConversionSingleRequiredObjectRaw]   `json:"single_required"`
}

func (d ConversionRaw) ToPlain() Conversion {
	return Conversion{
		MultipleOptional: escapeSlice(mapToPlain[ConversionMultipleOptionalNested](d.MultipleOptional.ValueMultiple())),
		MultipleRequired: mapToPlain[ConversionMultipleRequiredNested](d.MultipleRequired.ValueMultiple()),
		SingleOptional:   escapeValue(d.SingleOptional.PlainSingle().ToPlain()),
		SingleRequired:   d.SingleRequired.ValueSingle().ToPlain(),
	}
}

type ConversionMultipleOptionalNested struct {
	Age  int32                                `json:"age"`
	Name ConversionMultipleOptionalNameObject `json:"name"`
}

func (d ConversionMultipleOptionalNested) ToRaw() ConversionMultipleOptionalNestedRaw {
	return ConversionMultipleOptionalNestedRaw{
		Age:  elastic.FromSingle(d.Age),
		Name: elastic.FromSingle(d.Name.ToRaw()),
	}
}

type ConversionMultipleOptionalNestedRaw struct {
	Age  elastic.Elastic[int32]                                   `json:"age"`
	Name elastic.Elastic[ConversionMultipleOptionalNameObjectRaw] `json:"name"`
}

func (d ConversionMultipleOptionalNestedRaw) ToPlain() ConversionMultipleOptionalNested {
	return ConversionMultipleOptionalNested{
		Age:  d.Age.ValueSingle(),
		Name: d.Name.ValueSingle().ToPlain(),
	}
}

type ConversionMultipleOptionalNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func (d ConversionMultipleOptionalNameObject) ToRaw() ConversionMultipleOptionalNameObjectRaw {
	return ConversionMultipleOptionalNameObjectRaw{
		First: elastic.FromSingle(d.First),
		Last:  elastic.FromSingle(d.Last),
	}
}

type ConversionMultipleOptionalNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

func (d ConversionMultipleOptionalNameObjectRaw) ToPlain() ConversionMultipleOptionalNameObject {
	return ConversionMultipleOptionalNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}

type ConversionMultipleRequiredNested struct {
	Age  int32                                `json:"age"`
	Name ConversionMultipleRequiredNameObject `json:"name"`
}

func (d ConversionMultipleRequiredNested) ToRaw() ConversionMultipleRequiredNestedRaw {
	return ConversionMultipleRequiredNestedRaw{
		Age:  elastic.FromSingle(d.Age),
		Name: elastic.FromSingle(d.Name.ToRaw()),
	}
}

type ConversionMultipleRequiredNestedRaw struct {
	Age  elastic.Elastic[int32]                                   `json:"age"`
	Name elastic.Elastic[ConversionMultipleRequiredNameObjectRaw] `json:"name"`
}

func (d ConversionMultipleRequiredNestedRaw) ToPlain() ConversionMultipleRequiredNested {
	return ConversionMultipleRequiredNested{
		Age:  d.Age.ValueSingle(),
		Name: d.Name.ValueSingle().ToPlain(),
	}
}

type ConversionMultipleRequiredNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func (d ConversionMultipleRequiredNameObject) ToRaw() ConversionMultipleRequiredNameObjectRaw {
	return ConversionMultipleRequiredNameObjectRaw{
		First: elastic.FromSingle(d.First),
		Last:  elastic.FromSingle(d.Last),
	}
}

type ConversionMultipleRequiredNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

func (d ConversionMultipleRequiredNameObjectRaw) ToPlain() ConversionMultipleRequiredNameObject {
	return ConversionMultipleRequiredNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}

type ConversionSingleOptionalObject struct {
	Age  int32                              `json:"age"`
	Name ConversionSingleOptionalNameObject `json:"name"`
}

func (d ConversionSingleOptionalObject) ToRaw() ConversionSingleOptionalObjectRaw {
	return ConversionSingleOptionalObjectRaw{
		Age:  elastic.FromSingle(d.Age),
		Name: elastic.FromSingle(d.Name.ToRaw()),
	}
}

type ConversionSingleOptionalObjectRaw struct {
	Age  elastic.Elastic[int32]                                 `json:"age"`
	Name elastic.Elastic[ConversionSingleOptionalNameObjectRaw] `json:"name"`
}

func (d ConversionSingleOptionalObjectRaw) ToPlain() ConversionSingleOptionalObject {
	return ConversionSingleOptionalObject{
		Age:  d.Age.ValueSingle(),
		Name: d.Name.ValueSingle().ToPlain(),
	}
}

type ConversionSingleOptionalNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func (d ConversionSingleOptionalNameObject) ToRaw() ConversionSingleOptionalNameObjectRaw {
	return ConversionSingleOptionalNameObjectRaw{
		First: elastic.FromSingle(d.First),
		Last:  elastic.FromSingle(d.Last),
	}
}

type ConversionSingleOptionalNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

func (d ConversionSingleOptionalNameObjectRaw) ToPlain() ConversionSingleOptionalNameObject {
	return ConversionSingleOptionalNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}

type ConversionSingleRequiredObject struct {
	Age  int32                              `json:"age"`
	Name ConversionSingleRequiredNameObject `json:"name"`
}

func (d ConversionSingleRequiredObject) ToRaw() ConversionSingleRequiredObjectRaw {
	return ConversionSingleRequiredObjectRaw{
		Age:  elastic.FromSingle(d.Age),
		Name: elastic.FromSingle(d.Name.ToRaw()),
	}
}

type ConversionSingleRequiredObjectRaw struct {
	Age  elastic.Elastic[int32]                                 `json:"age"`
	Name elastic.Elastic[ConversionSingleRequiredNameObjectRaw] `json:"name"`
}

func (d ConversionSingleRequiredObjectRaw) ToPlain() ConversionSingleRequiredObject {
	return ConversionSingleRequiredObject{
		Age:  d.Age.ValueSingle(),
		Name: d.Name.ValueSingle().ToPlain(),
	}
}

type ConversionSingleRequiredNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func (d ConversionSingleRequiredNameObject) ToRaw() ConversionSingleRequiredNameObjectRaw {
	return ConversionSingleRequiredNameObjectRaw{
		First: elastic.FromSingle(d.First),
		Last:  elastic.FromSingle(d.Last),
	}
}

type ConversionSingleRequiredNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

func (d ConversionSingleRequiredNameObjectRaw) ToPlain() ConversionSingleRequiredNameObject {
	return ConversionSingleRequiredNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}
