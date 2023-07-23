package generator

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/ngicks/estype/gentypehelper"
)

func nameFromTag(s string) string {
	before, _, _ := strings.Cut(s, ",")
	return before
}

func generateAdditionalPropMarshalJSON(ctx *generatorContext, tyId typeId, fields []structField, isUndSerde bool) {
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
				Id(tyId.Id),
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
						Qual(gentypehelperQual, gentypehelper.IdGetBuf).
						Call(),
					jen.Defer().Qual(gentypehelperQual, gentypehelper.IdPutBuf).Call(jen.Id("buf")),
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

func generateAdditionalPropUnmarshalJSON(ctx *generatorContext, tyId typeId, fields []structField) {
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
				Op("*").
				Id(tyId.Id),
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
							for _, field := range fields {
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
