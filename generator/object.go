package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

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

func ObjectLike(ctx *GeneratorContext, dryRun bool) (plain, raw TypeId) {
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

	// Nested is not documented that it would inherit parent's dynamic if not explicitly set.
	// For now, keep inheritance. See behavior on real Elasticsearch instances and decide.
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

			var fieldTypeId TypeId
			var isObjectLike bool
			if mapping.IsObjectLike(propChild) {
				isObjectLike = true

				fieldHigh, fieldRaw := ObjectLike(nextCtx, true)
				if !opt.IsRaw {
					fieldTypeId = fieldHigh
				} else {
					fieldTypeId = fieldRaw
				}
			} else {
				fieldTypeId = Field(nextCtx, true)
			}

			fieldTypeId = opt.Mapper(fieldTypeId)
			fields = append(fields, structField{
				Name:         pascalCase(propFieldName),
				IsObjectLike: isObjectLike,
				Opt:          opt.TypeIdRenderOption(nextCtx),
				Stmt:         fieldTypeId.Render(opt.TypeIdRenderOption(nextCtx)),
				TypeId:       fieldTypeId,
				Tag:          map[string]string{"json": propFieldName},
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
		field := next.Latter
		nextCtx := ctx.next(fieldName, field, dynamic)

		if mapping.IsObjectLike(field) {
			_, _ = ObjectLike(nextCtx, false)
		} else {
			_ = Field(nextCtx, false)
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

func generateToRaw(ctx *GeneratorContext, plain, raw TypeId, plainFields, rawFields []structField, strict bool) {
	ctx.file.Func().
		Params(jen.Id("d").Id(plain.Id)).
		Id("ToRaw").
		Params().
		Id(raw.Id).Block(
		jen.Return(
			jen.Id(raw.Id).CustomFunc(
				jen.Options{
					Open:      "{",
					Close:     "}",
					Separator: ",",
					Multi:     true,
				},
				func(g *jen.Group) {
					for idx, field := range rawFields {
						stmt := jen.
							Id(field.Name).
							Op(":")

						constructorFnName := "From"
						if plainFields[idx].Opt.IsSingle() {
							constructorFnName += "Single"
						} else {
							constructorFnName += "Multiple"
						}
						if plainFields[idx].Opt.IsOptional() {
							constructorFnName += "Pointer"
						}
						stmt = stmt.
							Qual(elasticTypeQual, constructorFnName)

						var input *jen.Statement
						if !plainFields[idx].IsObjectLike {
							input = jen.Id("d").Dot(field.Name)
						} else {
							if plainFields[idx].Opt.IsSingle() {
								input = jen.Id("d").Dot(field.Name).Dot("ToRaw").Call()
							} else {
								retType := func(enableOpt bool) *jen.Statement {
									retType := new(jen.Statement)
									if enableOpt && plainFields[idx].Opt.IsOptional() {
										retType = retType.Op("*")
									}
									retType = retType.Index()
									retType = retType.Id(field.TypeId.TypeParam[0].Id)
									return retType
								}
								input = jen.
									Func().             // func
									Params().           // ()
									Add(retType(true)). // T
									Block(              // {
										jen.
											Id("out").
											Op(":=").
											Id("make").
											Call(
												retType(false),
												jen.Id("len").Call(jen.Id("d").Dot(field.Name)),
											),
										jen.
											For(
												jen.
													List(jen.Id("i"), jen.Id("v")).
													Op(":=").
													Range().
													Id("d").
													Dot(field.Name),
											).
											Block(
												jen.
													Id("out").
													Index(jen.Id("i")).
													Op("=").
													Id("v").
													Dot("ToRaw").
													Call(),
											),
										jen.ReturnFunc(func(g *jen.Group) {
											stmt := new(jen.Statement)
											if plainFields[idx].Opt.IsOptional() {
												stmt = stmt.Op("&")
											}
											g.Add(stmt.Id("out"))
										}),
									). // }
									Call()
							}
						}

						stmt = stmt.Call(input)

						g.Add(stmt)
					}
				},
			),
		),
	)
}

func generateToPlain(ctx *GeneratorContext, plain, raw TypeId, plainFields, rawFields []structField, strict bool) {
	ctx.file.Func().
		Params(jen.Id("d").Id(raw.Id)).
		Id("ToPlain").
		Params().
		Id(plain.Id).Block(
		jen.Return(
			jen.Id(plain.Id).CustomFunc(
				jen.Options{
					Open:      "{",
					Close:     "}",
					Separator: ",",
					Multi:     true,
				},
				func(g *jen.Group) {
					for idx, field := range plainFields {
						stmt := jen.
							Id(field.Name).
							Op(":")

						var constructorFnName string
						if field.Opt.IsOptional() {
							constructorFnName += "Plain"
						} else {
							constructorFnName += "Value"
						}
						if field.Opt.IsSingle() {
							constructorFnName += "Single"
						} else {
							constructorFnName += "Multiple"
						}

						if !field.IsObjectLike {
							stmt = stmt.Id("d").Dot(field.Name).Dot(constructorFnName).Call()
						} else {
							if field.Opt.IsSingle() {
								stmt = stmt.Id("d").Dot(field.Name).Dot(constructorFnName).Call().Dot("ToPlain").Call()
							} else {
								retType := func(enableOpt bool) *jen.Statement {
									retType := new(jen.Statement)
									if enableOpt && plainFields[idx].Opt.IsOptional() {
										retType = retType.Op("*")
									}
									retType = retType.Index().Id(field.TypeId.Id)
									return retType
								}
								stmt = stmt.
									Func().   // func
									Params(). // ()
									Add(      // T
										retType(true),
									).
									Block( // {
										jen.
											Id("values").
											Op(":=").
											Id("d").
											Dot(field.Name).
											Dot("ValueMultiple").
											Call(),
										jen.
											Id("out").
											Op(":=").
											Id("make").
											Call(
												retType(false),
												jen.Id("len").Call(jen.Id("values")),
											),
										jen.
											For(
												jen.
													List(jen.Id("i"), jen.Id("v")).
													Op(":=").
													Range().
													Id("values"),
											).
											Block(
												jen.
													Id("out").
													Index(jen.Id("i")).
													Op("=").
													Id("v").
													Dot("ToPlain").
													Call(),
											),
										jen.ReturnFunc(func(g *jen.Group) {
											stmt := new(jen.Statement)
											if plainFields[idx].Opt.IsOptional() {
												stmt = stmt.Op("&")
											}
											g.Add(stmt.Id("out"))
										}),
									).     // }
									Call() // ()
							}
						}

						g.Add(stmt)
					}
				},
			),
		),
	)
}

func generateBufPool(ctx *GeneratorContext) {
	if ctx.globalState.bufPoolGenerated {
		return
	}

	ctx.globalState.bufPoolGenerated = true
	ctx.file.Commentf("%s is pool of buffers. Implementations of MarshalJSON and UnmarshalJSON will use this.", bufPoolId)
	ctx.file.Var().Id(bufPoolId).Op("=").Op("&").Qual("sync", "Pool").Values(
		jen.Id("New").Op(":").Func().Params().Any().Block(
			jen.Return().New(jen.Qual("bytes", "Buffer")),
		),
	)
}
func generateAdditionalPropMarshalJSON(ctx *GeneratorContext, typeId TypeId, fields []structField, isUndSerde bool) {
	var marshalerStmt *jen.Statement
	if !isUndSerde {
		marshalerStmt = jen.Qual("encoding/json", "Marshal")
	} else {
		marshalerStmt = jen.Qual("github.com/ngicks/und/serde", "Marshal")
	}
	ctx.file.
		Comment(
			"// MarshalJSON implements json.Marshaler\n"+
				"// so that both known fields and additional properties are marshaled into a same JSON object.\n"+
				"//\n"+
				"// The presence of this implementation indicates the dynamic field for this object are\n"+
				"// defined to be other than \"strict\" in its mapping.json.",
		).
		Line().
		Func().
		Params(
			jen.
				Id("d").
				Id(typeId.Id),
		).
		Id("MarshalJSON").
		Params().
		Params(
			jen.Index().Byte(),
			jen.Error(),
		).
		BlockFunc(
			func(g *jen.Group) {
				stmts := make([]*jen.Statement, 0)
				stmts = append(stmts, []*jen.Statement{
					jen.
						Id("buf").
						Op(":=").
						Id(bufPoolId).
						Dot("Get").
						Call().
						Op(".").
						Parens(jen.Op("*").Qual("bytes", "Buffer")),
					jen.Defer().Id(bufPoolId).Dot("Put").Call(jen.Id("buf")),
					jen.Id("buf").Dot("Reset").Call(),
					jen.
						Var().
						Defs(
							jen.Id("bin").Index().Byte(),
							jen.Err().Error(),
						),
					jen.
						Id("buf").
						Dot("WriteByte").
						Call(jen.Id(`'{'`)),
				}...)
				buf := new(bytes.Buffer)
				for _, field := range fields {
					json.HTMLEscape(buf, []byte(nameFromTag(field.Tag["json"])))
					escapedName := buf.String()
					buf.Reset()
					stmts = append(stmts, []*jen.Statement{
						jen.
							Id("buf").
							Dot("WriteString").
							Call(jen.Lit(`"` + escapedName + `":`)),
						jen.
							List(jen.Id("bin"), jen.Err()).
							Op("=").
							Add(marshalerStmt).
							Call(jen.Id("d").Dot(field.Name)),

						jen.If(jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Nil(), jen.Err()),
						),
						jen.Id("buf").Dot("Write").Call(jen.Id("bin")),
						jen.Id("buf").Dot("WriteByte").Call(jen.Id(`','`)),
					}...)
				}
				stmts = append(stmts, []*jen.Statement{
					jen.Id("keys").Op(":=").Make(jen.Index().String(), jen.Lit(0), jen.Len(jen.Id("d").Dot(additionalPropId))),
					jen.For(jen.Id("k").Op(":=").Range().Id("d").Dot(additionalPropId)).Block(
						jen.Id("keys").Op("=").Append(jen.Id("keys"), jen.Id("k")),
					),
					jen.Qual("sort", "Strings").Call(jen.Id("keys")),
					jen.For(jen.List(jen.Id("_"), jen.Id("key")).Op(":=").Range().Id("keys")).Block(
						jen.List(jen.Id("bin"), jen.Err()).Op("=").Add(marshalerStmt).Call(jen.Id("d").Dot(additionalPropId).Index(jen.Id("key"))),
						jen.If(jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Nil(), jen.Err()),
						),
						jen.Id("buf").Dot("WriteByte").Call(jen.Id(`'"'`)),
						jen.Id("buf").Dot("WriteString").Call(jen.Id("key")),
						jen.Id("buf").Dot("WriteString").Call(jen.Lit(`":`)),
						jen.Id("buf").Dot("Write").Call(jen.Id("bin")),
						jen.Id("buf").Dot("WriteByte").Call(jen.Id(`','`)),
					),
					jen.If(jen.Id("buf").Dot("Len").Call().Op(">").Lit(1)).Block(
						jen.Id("buf").Dot("Truncate").Call(jen.Id("buf").Dot("Len").Call().Op("-").Lit(1)),
					),
					jen.Id("buf").Dot("WriteByte").Call(jen.Id(`'}'`)),
					jen.Return(
						jen.Append(
							jen.Index().Byte().Values(),
							jen.Id("buf").Dot("Bytes").Call().Op("..."),
						),
						jen.Nil(),
					),
				}...)

				for _, stmt := range stmts {
					g.Add(stmt)
				}
			},
		)
}

func nameFromTag(s string) string {
	before, _, _ := strings.Cut(s, ",")
	return before
}

func generateAdditionalPropUnmarshalJSON(ctx *GeneratorContext, plain TypeId, plainFields []structField) {
	ctx.file.
		Commentf(
			"// UnmarshalJSON implements json.Unmarshaler\n"+
				"// to add the special handling rule where\n"+
				"// additional fields in the input JSON are stored into the %s field\n"+
				"//\n"+
				"// The presence of this implementation indicates the dynamic field for this object are\n"+
				"// defined to be other than \"strict\" in its mapping.json.",
			additionalPropId,
		).
		Line().
		Func().
		Params(
			jen.
				Id("d").
				Id(plain.Id),
		).
		Id("UnmarshalJSON").
		Params(jen.Id("data").Index().Byte()).
		Error().
		BlockFunc(func(g *jen.Group) {
			stmts := make([]*jen.Statement, 0)

			stmts = append(stmts, []*jen.Statement{
				jen.
					Id("dec").
					Op(":=").
					Qual("encoding/json", "NewDecoder").
					Call(
						jen.
							Qual("bytes", "NewBuffer").
							Call(jen.Id("data")),
					),
				jen.
					List(jen.Id("token"), jen.Err()).
					Op(":=").
					Id("dec").
					Dot("Token").
					Call(),
				jen.If(jen.Err().Op("!=").Nil()).Block(
					jen.Return(jen.Err()),
				),
				jen.If(jen.Id("token").Op("!=").Qual("encoding/json", "Delim").Parens(jen.Id(`'{'`))).Block(
					jen.Return(
						jen.
							Qual("fmt", "Errorf").
							Call(
								jen.Lit(
									"unknown token. Assuming the input is a JSON object,"+
										" but received wrong delim = %s",
								),
								jen.Id("token"),
							),
					),
				),
				jen.Id("firstWriteToAdditionalProp").Op(":=").Lit(true),
				jen.For().BlockFunc(func(g *jen.Group) {
					stmts := make([]*jen.Statement, 0)

					stmts = append(stmts, []*jen.Statement{
						jen.
							List(jen.Id("token"), jen.Err()).
							Op(":=").
							Id("dec").
							Dot("Token").
							Call(),
						jen.If(jen.Err().Op("!=").Nil()).Block(
							jen.If(jen.Qual("errors", "Is").Call(jen.Err(), jen.Qual("io", "EOF"))).Block(
								jen.Break(),
							),
							jen.Return(jen.Err()),
						),
						jen.Switch(jen.Id("token")).BlockFunc(func(g *jen.Group) {
							for _, field := range plainFields {
								g.Add(
									jen.Case(
										jen.Lit(nameFromTag(field.Tag["json"])),
									),
								)
								g.Add(
									jen.
										Err().
										Op("=").
										Id("dec").
										Dot("Decode").
										Call(
											jen.Op("&").Id("d").Dot(field.Name),
										),
								)
							}
							g.Add(jen.Default())
							g.Add(
								jen.If(
									jen.List(jen.Id("key"), jen.Id("ok")).
										Op(":=").
										Id("token").
										Op(".").
										Parens(jen.String()).
										Op(";").
										Id("ok"),
								).Block(
									jen.Var().Id("o").Any(),
									jen.Err().Op("=").Id("dec").Dot("Decode").Call(jen.Op("&").Id("o")),
									jen.If(jen.Err().Op("!=").Nil()).Block(
										jen.Return(jen.Err()),
									),
									jen.Comment("// map re-initialization is deferred until at least a successful decode."),
									jen.If(jen.Id("firstWriteToAdditionalProp")).Block(
										jen.Id("firstWriteToAdditionalProp").Op("=").False(),
										jen.Id("d").Dot(additionalPropId).Op("=").Make(jen.Map(jen.String()).Any()),
									),
									jen.Id("d").Dot(additionalPropId).Index(jen.Id("key")).Op("=").Id("o"),
								),
							)
						}),
						jen.If(jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Err()),
						),
					}...)

					for _, stmt := range stmts {
						g.Add(stmt)
					}
				}),
				jen.Return(jen.Nil()),
			}...)

			for _, stmt := range stmts {
				g.Add(stmt)
			}
		})
}
