package test

import (
	gentypehelper "github.com/ngicks/estype/gentypehelper"
	elastic "github.com/ngicks/und/elastic"
)

type Conversion struct {
	MultipleOptional *[]ConversionMultipleOptionalNested `json:"multiple_optional"`
	MultipleRequired []ConversionMultipleRequiredNested  `json:"multiple_required"`
	SingleOptional   *ConversionSingleOptionalObject     `json:"single_optional"`
	SingleRequired   ConversionSingleRequiredObject      `json:"single_required"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d Conversion) ToRaw() ConversionRaw {
	return ConversionRaw{
		MultipleOptional: gentypehelper.MapPlainMultipleOptionalToRawElastic[ConversionMultipleOptionalNestedRaw](d.MultipleOptional),
		MultipleRequired: gentypehelper.MapPlainMultipleToRawElastic[ConversionMultipleRequiredNestedRaw](d.MultipleRequired),
		SingleOptional:   gentypehelper.MapPlainOptionalToRawElastic[ConversionSingleOptionalObjectRaw](d.SingleOptional),
		SingleRequired:   gentypehelper.MapPlainToRawElastic[ConversionSingleRequiredObjectRaw](d.SingleRequired),
	}
}

type ConversionRaw struct {
	MultipleOptional elastic.Elastic[ConversionMultipleOptionalNestedRaw] `json:"multiple_optional"`
	MultipleRequired elastic.Elastic[ConversionMultipleRequiredNestedRaw] `json:"multiple_required"`
	SingleOptional   elastic.Elastic[ConversionSingleOptionalObjectRaw]   `json:"single_optional"`
	SingleRequired   elastic.Elastic[ConversionSingleRequiredObjectRaw]   `json:"single_required"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionRaw) ToPlain() Conversion {
	return Conversion{
		MultipleOptional: gentypehelper.MapElasticToPlainMultipleOptional[ConversionMultipleOptionalNested](d.MultipleOptional),
		MultipleRequired: gentypehelper.MapElasticToPlainMultiple[ConversionMultipleRequiredNested](d.MultipleRequired),
		SingleOptional:   gentypehelper.MapElasticToPlainSingleOptional[ConversionSingleOptionalObject](d.SingleOptional),
		SingleRequired:   gentypehelper.MapElasticToPlainSingle[ConversionSingleRequiredObject](d.SingleRequired),
	}
}

type ConversionMultipleOptionalNested struct {
	Age  int32                                `json:"age"`
	Name ConversionMultipleOptionalNameObject `json:"name"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionMultipleOptionalNested) ToRaw() ConversionMultipleOptionalNestedRaw {
	return ConversionMultipleOptionalNestedRaw{
		Age:  gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainToRawElastic[ConversionMultipleOptionalNameObjectRaw](d.Name),
	}
}

type ConversionMultipleOptionalNestedRaw struct {
	Age  elastic.Elastic[int32]                                   `json:"age"`
	Name elastic.Elastic[ConversionMultipleOptionalNameObjectRaw] `json:"name"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionMultipleOptionalNestedRaw) ToPlain() ConversionMultipleOptionalNested {
	return ConversionMultipleOptionalNested{
		Age:  d.Age.ValueSingle(),
		Name: gentypehelper.MapElasticToPlainSingle[ConversionMultipleOptionalNameObject](d.Name),
	}
}

type ConversionMultipleOptionalNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionMultipleOptionalNameObject) ToRaw() ConversionMultipleOptionalNameObjectRaw {
	return ConversionMultipleOptionalNameObjectRaw{
		First: gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:  gentypehelper.MapSingleValueToElastic[string](d.Last),
	}
}

type ConversionMultipleOptionalNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
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

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionMultipleRequiredNested) ToRaw() ConversionMultipleRequiredNestedRaw {
	return ConversionMultipleRequiredNestedRaw{
		Age:  gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainToRawElastic[ConversionMultipleRequiredNameObjectRaw](d.Name),
	}
}

type ConversionMultipleRequiredNestedRaw struct {
	Age  elastic.Elastic[int32]                                   `json:"age"`
	Name elastic.Elastic[ConversionMultipleRequiredNameObjectRaw] `json:"name"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionMultipleRequiredNestedRaw) ToPlain() ConversionMultipleRequiredNested {
	return ConversionMultipleRequiredNested{
		Age:  d.Age.ValueSingle(),
		Name: gentypehelper.MapElasticToPlainSingle[ConversionMultipleRequiredNameObject](d.Name),
	}
}

type ConversionMultipleRequiredNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionMultipleRequiredNameObject) ToRaw() ConversionMultipleRequiredNameObjectRaw {
	return ConversionMultipleRequiredNameObjectRaw{
		First: gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:  gentypehelper.MapSingleValueToElastic[string](d.Last),
	}
}

type ConversionMultipleRequiredNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
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

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionSingleOptionalObject) ToRaw() ConversionSingleOptionalObjectRaw {
	return ConversionSingleOptionalObjectRaw{
		Age:  gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainToRawElastic[ConversionSingleOptionalNameObjectRaw](d.Name),
	}
}

type ConversionSingleOptionalObjectRaw struct {
	Age  elastic.Elastic[int32]                                 `json:"age"`
	Name elastic.Elastic[ConversionSingleOptionalNameObjectRaw] `json:"name"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionSingleOptionalObjectRaw) ToPlain() ConversionSingleOptionalObject {
	return ConversionSingleOptionalObject{
		Age:  d.Age.ValueSingle(),
		Name: gentypehelper.MapElasticToPlainSingle[ConversionSingleOptionalNameObject](d.Name),
	}
}

type ConversionSingleOptionalNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionSingleOptionalNameObject) ToRaw() ConversionSingleOptionalNameObjectRaw {
	return ConversionSingleOptionalNameObjectRaw{
		First: gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:  gentypehelper.MapSingleValueToElastic[string](d.Last),
	}
}

type ConversionSingleOptionalNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
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

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionSingleRequiredObject) ToRaw() ConversionSingleRequiredObjectRaw {
	return ConversionSingleRequiredObjectRaw{
		Age:  gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainToRawElastic[ConversionSingleRequiredNameObjectRaw](d.Name),
	}
}

type ConversionSingleRequiredObjectRaw struct {
	Age  elastic.Elastic[int32]                                 `json:"age"`
	Name elastic.Elastic[ConversionSingleRequiredNameObjectRaw] `json:"name"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionSingleRequiredObjectRaw) ToPlain() ConversionSingleRequiredObject {
	return ConversionSingleRequiredObject{
		Age:  d.Age.ValueSingle(),
		Name: gentypehelper.MapElasticToPlainSingle[ConversionSingleRequiredNameObject](d.Name),
	}
}

type ConversionSingleRequiredNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionSingleRequiredNameObject) ToRaw() ConversionSingleRequiredNameObjectRaw {
	return ConversionSingleRequiredNameObjectRaw{
		First: gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:  gentypehelper.MapSingleValueToElastic[string](d.Last),
	}
}

type ConversionSingleRequiredNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d ConversionSingleRequiredNameObjectRaw) ToPlain() ConversionSingleRequiredNameObject {
	return ConversionSingleRequiredNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}
