package generator

import (
	"fmt"
	"sort"

	"github.com/dave/jennifer/jen"
	"github.com/ngicks/estype/spec/mapping"
	"github.com/ngicks/go-iterator/iterator"
	"github.com/ngicks/go-iterator/iterator/adapter"
	"github.com/ngicks/und/option"
)

const (
	elasticTypeQual       = "github.com/ngicks/und/elastic"
	undefinedableTypeQual = "github.com/ngicks/und/undefinedable"
	jsonfieldTypeQual     = "github.com/ngicks/und/jsonfield"
	utilTypeQual          = "github.com/ngicks/estype/util"
	additionalPropId      = "AdditionalProperties_"
)

var (
	elasticTypeId       = typeId{Id: "Elastic", Qualifier: elasticTypeQual}
	undefinedableTypeId = typeId{Id: "Undefinedable", Qualifier: undefinedableTypeQual}
	jsonfieldTypeId     = typeId{Id: "JsonField", Qualifier: jsonfieldTypeQual}
)

type structField struct {
	Name string
	// raw name for this field.
	// This would not be same as Name field wehen the field contains non-ident string.
	// Suffixing field name with `Raw` should be avoided in this struct since it may confuse people
	// as they assume Raw is related to Raw variant of generated code.
	NameUnprocessed string
	IsObjectLike    bool // nested or object field data type.
	Opt             typeIdRenderOption
	Stmt            *jen.Statement
	TypeId          typeId
	Tag             map[string]string
}

func genObjectLike(ctx *generatorContext, dryRun bool) (plain, raw typeId) {
	plain = typeId{
		Id: ctx.getTypeName(),
	}
	raw = typeId{
		Id: plain.Id + "Raw",
	}

	if dryRun {
		return plain, raw
	}

	var props map[string]mapping.Property
	var dynamic option.Option[mapping.DynamicMapping]
	var strict bool
	switch x := ctx.localState.prop.Val.(type) {
	case mapping.ObjectProperty:
		props = x.Properties.Value()
		dynamic = x.Dynamic.Option
	case mapping.NestedProperty:
		props = x.Properties.Value()
		dynamic = x.Dynamic.Option
	default:
		panic(fmt.Errorf("unknown type: %T", ctx.localState.prop))
	}

	// While the document is not saying that, nested also inherit dynamic prop from its parent.
	switch dynamic.Or(ctx.localState.dynamic).Value() {
	case mapping.Strict:
		strict = true
	default:
		strict = false
	}

	declMap := map[string][]structField{}

	type renderOpt struct {
		Id                 typeId
		IsRaw              bool
		TypeIdRenderOption func(ctx *generatorContext) typeIdRenderOption
		Mapper             func(i typeId) typeId
	}

	for _, opt := range []renderOpt{
		{
			Id:                 plain,
			TypeIdRenderOption: func(ctx *generatorContext) typeIdRenderOption { return ctx },
			Mapper:             func(i typeId) typeId { return i },
		},
		{
			Id:                 raw,
			IsRaw:              true,
			TypeIdRenderOption: func(_ *generatorContext) typeIdRenderOption { return newSimpleRenderOption(false, true) },
			Mapper: func(i typeId) typeId {
				var tyId typeId
				if i.DisallowArray {
					if i.DisallowNull {
						tyId = undefinedableTypeId
					} else {
						tyId = jsonfieldTypeId
					}
				} else {
					// The type may only allow a single value
					// however it may also allow an array containing only a single element.
					tyId = elasticTypeId
				}
				tyId.TypeParam = []typeId{i}
				return tyId
			},
		},
	} {
		fields := make([]structField, 0)

		iter := createPropsIter(props)
		for next, ok := iter.Next(); ok; next, ok = iter.Next() {
			propFieldName := next.Former
			propChild := next.Latter
			nextCtx := ctx.next(propFieldName, propChild, dynamic)

			var fieldTypeId typeId
			var isObjectLike bool
			if mapping.IsObjectLike(propChild) {
				isObjectLike = true

				fieldPlain, fieldRaw := genObjectLike(nextCtx, true)
				if !opt.IsRaw {
					fieldTypeId = fieldPlain
				} else {
					fieldTypeId = fieldRaw
				}
			} else {
				fieldTypeId = genField(nextCtx, true)
			}

			mappedFieldTypeId := opt.Mapper(fieldTypeId)
			var omitemptyOpt string
			if mappedFieldTypeId.MustOmit(nextCtx) {
				omitemptyOpt = ",omitempty"
			}

			fields = append(fields, structField{
				Name:            pascalCase(escapeNonId(propFieldName)),
				NameUnprocessed: pascalCase(propFieldName),
				IsObjectLike:    isObjectLike,
				Opt:             opt.TypeIdRenderOption(nextCtx),
				Stmt:            mappedFieldTypeId.Render(opt.TypeIdRenderOption(nextCtx)),
				TypeId:          mappedFieldTypeId,
				Tag:             map[string]string{"json": propFieldName + omitemptyOpt},
			})
		}

		declMap[opt.Id.Id] = fields
	}

	ctx.file.Type().Id(plain.Id).StructFunc(func(g *jen.Group) {
		for _, field := range declMap[plain.Id] {
			g.Add(jen.Id(field.Name).Add(field.Stmt).Tag(field.Tag))
		}
		if !strict {
			g.Add(jen.Id(additionalPropId).Id(anyMap))
		}
	})

	generateToRaw(ctx, plain, raw, declMap[plain.Id], declMap[raw.Id], strict)

	if !strict {
		generateAdditionalPropMarshalJSON(ctx, plain, declMap[plain.Id], false)
		generateAdditionalPropUnmarshalJSON(ctx, plain, declMap[plain.Id])
	}

	ctx.file.Type().Id(raw.Id).StructFunc(func(g *jen.Group) {
		for _, field := range declMap[raw.Id] {
			g.Add(jen.Id(field.Name).Add(field.Stmt).Tag(field.Tag))
		}
		if !strict {
			g.Add(jen.Id(additionalPropId).Id(anyMap))
		}
	})

	generateToPlain(ctx, plain, raw, declMap[plain.Id], declMap[raw.Id], strict)
	if !strict {
		generateAdditionalPropMarshalJSON(ctx, raw, declMap[raw.Id], true)
		generateAdditionalPropUnmarshalJSON(ctx, raw, declMap[raw.Id])
	}

	// render subtypes.
	iter := createPropsIter(props)
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		fieldName := next.Former
		prop := next.Latter
		nextCtx := ctx.next(fieldName, prop, dynamic)

		if mapping.IsObjectLike(prop) {
			_, _ = genObjectLike(nextCtx, false)
		} else {
			_ = genField(nextCtx, false)
		}
	}

	return plain, raw
}

func createPropsIter(props map[string]mapping.Property) iterator.Iterator[adapter.Zipped[string, mapping.Property]] {
	return iterator.FromMap(
		props,
		func(keys []string) []string {
			sort.Strings(keys)
			return keys
		},
	)
}
