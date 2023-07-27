package genestime

import (
	"github.com/dave/jennifer/jen"
	"github.com/ngicks/estype/fielddatatype/estime"
	"github.com/ngicks/und/option"
)

const (
	estimeQual                  = "github.com/ngicks/estype/fielddatatype/estime"
	timeParserConstructorFnName = "FromGoTimeLayoutUnsafe"
)

type GeneratorDef struct {
	TyName             string
	MultiLayout        estime.MultiLayout
	NumParser          estime.NumberParser
	MarshalToNumber    bool
	MarshalLayoutIndex uint
	Comment            option.Option[string] // Comment for type. If None, default comment will be inserted before type declaration.
}

func (d GeneratorDef) Gen(f *jen.File) {
	parserTyName := "parser" + d.TyName

	if d.Comment.IsSome() {
		if d.Comment.Value() != "" {
			f.Comment(d.Comment.Value())
		}
	} else {
		f.Commentf("// %s represents the date or the date_nanos mapping field type.", d.TyName)
		f.Comment("// It implements json.Unmarshaler so that it can be directly unmarshaled from\n" +
			"// all possible formats specified in corresponding `format` field.")
		f.Comment("//")
		f.Comment("// Allowed formats are:")
		f.Comment("//")
		for _, l := range d.MultiLayout.Clone() {
			f.Commentf("//  - %s", l)
		}
		if d.NumParser != "" {
			f.Commentf("//  - int as %s", d.NumParser)
		}
		f.Comment("//")
		f.Comment("// It also implements json.Marshaler. It will be marshaled into")
		if d.MarshalToNumber {
			f.Commentf("int which represents %s.", string(d.NumParser))
		} else {
			f.Commentf("string formatted in %s layout.", d.MultiLayout.Clone()[d.MarshalLayoutIndex])
		}
	}

	f.Type().Id(d.TyName).Qual("time", "Time")

	f.
		Var().
		Id(parserTyName).
		Op("=").
		Qual(estimeQual, timeParserConstructorFnName).
		Call(
			jen.
				Line().
				Index().String().CustomFunc(
				jen.Options{Open: "{", Close: "}", Multi: true, Separator: ","},
				func(g *jen.Group) {
					for _, v := range d.MultiLayout.Clone() {
						g.Add(jen.Lit(v))
					}
				},
			),
			jen.Line().Lit(string(d.NumParser)),
			jen.Line(),
		)

	f.Comment("// String implements fmt.Stringer")
	f.
		Func().                           // func
		Params(jen.Id("t").Id(d.TyName)). // (t <TyName>)
		Id("String").
		Params().
		Id("string").
		BlockFunc(func(g *jen.Group) {
			if d.MarshalToNumber {
				g.Add(
					jen.
						Return(
							jen.Qual("strconv", "FormatInt").Call(
								jen.Id(parserTyName).
									Dot("FormatNumber").
									Call(jen.Qual("time", "Time").Parens(jen.Id("t"))),
								jen.Lit(10),
							),
						),
				)
			} else {
				g.Add(
					jen.Return(
						jen.
							Id(parserTyName).
							Dot("FormatString").
							Call(
								jen.Qual("time", "Time").Parens(jen.Id("t")),
								// calling Lit w/ uint make it uint(0x0).
								// 0 is still untyped int so less verbose for eyes.
								jen.Lit(int(d.MarshalLayoutIndex)),
							),
					),
				)
			}
		})

	f.Comment("// MarshalJSON implements json.Marshaler")
	f.
		Func().                           // func
		Params(jen.Id("t").Id(d.TyName)). // (t <TyName>)
		Id("MarshalJSON").Params().       // MarshalJSON()
		Parens(
			jen.List(jen.Id("[]byte"), jen.Id("error")),
		). // ([]byte, error)
		BlockFunc(
			func(g *jen.Group) {
				stringerCall := jen.Id("t").Dot("String").Call()

				if !d.MarshalToNumber {
					stringerCall = jen.Lit("\"").Op("+").Add(stringerCall).Op("+").Lit("\"")
				}
				stringerCall = jen.Id("[]byte").Parens(stringerCall)
				g.Add(
					jen.Return(
						stringerCall,
						jen.Nil(),
					),
				)

			},
		)

	f.Comment("// UnmarshalJSON implements json.Unmarshaler")
	f.
		Func().                                                  // func
		Params(jen.Id("t").Op("*").Id(d.TyName)).                // (t <TyName>)
		Id("UnmarshalJSON").Params(jen.Id("data").Id("[]byte")). // UnmarshalJSON(data []byte)
		Parens(jen.Id("error")).                                 // error
		Block(
			jen.If(
				jen.
					String().
					Parens(
						jen.Id("data"),
					).
					Op("==").
					Lit("null"),
			).Block(
				jen.Return(jen.Nil()),
			),
			jen.List(
				jen.Id("tt"),
				jen.Id("err")).
				Op(":=").
				Id(parserTyName).
				Dot("ParseJson").
				Call(
					jen.Id("data"),
				),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Return(jen.Id("err")),
			),
			jen.Op("*").Id("t").Op("=").Id(d.TyName).Parens(jen.Id("tt")),
			jen.Return(jen.Nil()),
		)
}
