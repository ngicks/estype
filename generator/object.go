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
	elasticTypeQual  = "github.com/ngicks/und/elastic"
	utilTypeQual     = "github.com/ngicks/estype/util"
	additionalPropId = "AdditionalProperties_"
	bufPoolId        = "bufferPool"
)

type structField struct {
	Name         string
	IsObjectLike bool
	Opt          TypeIdRenderOption
	Stmt         *jen.Statement
	TypeId       TypeId
	Tag          map[string]string
}

func genObjectLike(ctx *GeneratorContext, dryRun bool) (plain, raw TypeId) {
	plain = TypeId{
		Id: ctx.getTypeName(),
	}
	raw = TypeId{
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
		Id                 TypeId
		IsRaw              bool
		TypeIdRenderOption func(ctx *GeneratorContext) TypeIdRenderOption
		Mapper             func(i TypeId) TypeId
	}

	for _, opt := range []renderOpt{
		{
			Id:                 plain,
			TypeIdRenderOption: func(ctx *GeneratorContext) TypeIdRenderOption { return ctx },
			Mapper:             func(i TypeId) TypeId { return i },
		},
		{
			Id:                 raw,
			IsRaw:              true,
			TypeIdRenderOption: func(_ *GeneratorContext) TypeIdRenderOption { return RenderOption(false, true) },
			Mapper: func(i TypeId) TypeId {
				return TypeId{Qualifier: elasticTypeQual, Id: "Elastic", TypeParam: []TypeId{i}}
			},
		},
	} {
		fields := make([]structField, 0)

		iter := createPropsIter(props)
		for next, ok := iter.Next(); ok; next, ok = iter.Next() {
			propFieldName := next.Former
			propChild := next.Latter
			nextCtx := ctx.next(propFieldName, propChild, dynamic)

			// generate helpers for later uses.
			if nextCtx.IsOptional() {
				if nextCtx.IsSingle() {
					generateEscapeValue(nextCtx)
				} else {
					generateEscapeSlice(nextCtx)
				}
			}
			if !nextCtx.IsSingle() {
				generateMapToPlain(nextCtx)
				if nextCtx.IsOptional() {
					generateMapToRawPointer(nextCtx)
				} else {
					generateMapToRaw(nextCtx)
				}
			}

			var fieldTypeId TypeId
			var isObjectLike bool
			if mapping.IsObjectLike(propChild) {
				isObjectLike = true

				fieldHigh, fieldRaw := genObjectLike(nextCtx, true)
				if !opt.IsRaw {
					fieldTypeId = fieldHigh
				} else {
					fieldTypeId = fieldRaw
				}
			} else {
				fieldTypeId = genField(nextCtx, true)
			}

			fieldTypeId = opt.Mapper(fieldTypeId)
			var omitemptyOpt string
			if fieldTypeId.NonWritable {
				omitemptyOpt = ",omitempty"
			}
			fields = append(fields, structField{
				Name:         pascalCase(propFieldName),
				IsObjectLike: isObjectLike,
				Opt:          opt.TypeIdRenderOption(nextCtx),
				Stmt:         fieldTypeId.Render(opt.TypeIdRenderOption(nextCtx)),
				TypeId:       fieldTypeId,
				Tag:          map[string]string{"json": propFieldName + omitemptyOpt},
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
		generateBufPool(ctx)
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
		generateBufPool(ctx)
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
