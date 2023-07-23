package generator

import (
	"github.com/dave/jennifer/jen"
	"github.com/ngicks/estype/gentypehelper"
)

const (
	gentypehelperQual = "github.com/ngicks/estype/gentypehelper"
)

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
						if plainFields[idx].TypeId.IgnoreInConversion() {
							continue
						}

						g.Add(
							jen.
								Id(field.Name).
								Op(":").
								Add(toRawStmt(plainFields[idx], field)),
						)
					}
					if !strict {
						g.Add(jen.Id(additionalPropId).Op(":").Id("d").Dot(additionalPropId))
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
						if field.TypeId.IgnoreInConversion() {
							continue
						}

						g.Add(
							jen.
								Id(field.Name).
								Op(":").
								Add(toPlainStmt(field, rawFields[idx])),
						)
					}
					if !strict {
						g.Add(jen.Id(additionalPropId).Op(":").Id("d").Dot(additionalPropId))
					}
				},
			),
		),
	)
}

func toRawStmt(plain, raw structField) *jen.Statement {
	fieldNameId := jen.Id("d").Dot(plain.Name)
	switch raw.TypeId.Qualifier {
	case jsonfieldTypeQual:
		// must be single at this point.
		if plain.TypeId.IsOptional(plain.Opt) {
			return jen.Qual(jsonfieldTypeQual, "FromPointer").Call(fieldNameId)
		} else {
			return jen.Qual(jsonfieldTypeQual, "Defined").Call(fieldNameId)
		}
	case undefinedableTypeQual:
		// must single and null is not acceptable.
		if plain.TypeId.IsOptional(plain.Opt) {
			return jen.Qual(undefinedableTypeQual, "FromPointer").Call(fieldNameId)
		} else {
			return jen.Qual(undefinedableTypeQual, "Defined").Call(fieldNameId)
		}
	case elasticTypeQual:
		return jen.
			Add(elasticMapper(plain)).
			Index(raw.TypeId.TypeParam[0].Render(RenderOption(false, true))).
			Call(fieldNameId)
	}
	panic("unknown")
}

func toPlainStmt(plain, raw structField) *jen.Statement {
	fieldNameId := jen.Id("d").Dot(raw.Name)

	switch raw.TypeId.Qualifier {
	case jsonfieldTypeQual:
		if plain.TypeId.IsOptional(plain.Opt) {
			return fieldNameId.Dot("Undefinedable").Dot("Value").Call().Dot("Plain").Call()
		} else {
			return fieldNameId.Dot("Value").Call()
		}
	case undefinedableTypeQual:
		if plain.TypeId.IsOptional(plain.Opt) {
			return fieldNameId.Dot("Plain").Call()
		} else {
			return fieldNameId.Dot("Value").Call()
		}
	case elasticTypeQual:
		if !plain.IsObjectLike {
			var valueMethodName string
			switch {
			case plain.TypeId.IsSingle(plain.Opt) && !plain.TypeId.IsOptional(plain.Opt):
				valueMethodName = "ValueSingle"
			case plain.TypeId.IsSingle(plain.Opt) && plain.TypeId.IsOptional(plain.Opt):
				valueMethodName = "PlainSingle"
			case !plain.TypeId.IsSingle(plain.Opt) && !plain.TypeId.IsOptional(plain.Opt):
				valueMethodName = "ValueMultiple"
			case !plain.TypeId.IsSingle(plain.Opt) && plain.TypeId.IsOptional(plain.Opt):
				return jen.
					Qual(gentypehelperQual, gentypehelper.IdMapElasticToMultipleValueOptional).
					Index(plain.TypeId.Render(RenderOption(false, true))).
					Call(fieldNameId)
			}
			return fieldNameId.Dot(valueMethodName).Call()
		} else {
			return jen.
				Add(plainMapper(plain)).
				Index(plain.TypeId.Render(RenderOption(false, true))).
				Call(fieldNameId)
		}
	}
	panic("unknown")
}

func elasticMapper(f structField) *jen.Statement {
	var fnName string
	if f.TypeId.DisallowNull && f.TypeId.IsOptional(f.Opt) {
		if f.TypeId.IsSingle(f.Opt) {
			fnName = gentypehelper.IdMapPlainPointerToUndefinedElastic
		} else {
			fnName = gentypehelper.IdMapPlainMultiplePointerToUndefinedElastic
		}

		return jen.Qual(gentypehelperQual, fnName)
	}

	if !f.IsObjectLike {
		switch {
		case f.TypeId.IsSingle(f.Opt) && !f.TypeId.IsOptional(f.Opt):
			fnName = gentypehelper.IdMapSingleValueToElastic
		case f.TypeId.IsSingle(f.Opt) && f.TypeId.IsOptional(f.Opt):
			fnName = gentypehelper.IdMapSingleOptionalValueToElastic
		case !f.TypeId.IsSingle(f.Opt) && !f.TypeId.IsOptional(f.Opt):
			fnName = gentypehelper.IdMapMultipleValueToElastic
		case !f.TypeId.IsSingle(f.Opt) && f.TypeId.IsOptional(f.Opt):
			fnName = gentypehelper.IdMapMultipleOptionalValueToElastic
		}
	} else {
		switch {
		case f.TypeId.IsSingle(f.Opt) && !f.TypeId.IsOptional(f.Opt):
			fnName = gentypehelper.IdMapPlainToRawElastic
		case f.TypeId.IsSingle(f.Opt) && f.TypeId.IsOptional(f.Opt):
			fnName = gentypehelper.IdMapPlainOptionalToRawElastic
		case !f.TypeId.IsSingle(f.Opt) && !f.TypeId.IsOptional(f.Opt):
			fnName = gentypehelper.IdMapPlainMultipleToRawElastic
		case !f.TypeId.IsSingle(f.Opt) && f.TypeId.IsOptional(f.Opt):
			fnName = gentypehelper.IdMapPlainMultipleOptionalToRawElastic
		}
	}

	return jen.Qual(gentypehelperQual, fnName)
}

func plainMapper(f structField) *jen.Statement {
	var fnName string
	switch {
	case f.TypeId.IsSingle(f.Opt) && !f.TypeId.IsOptional(f.Opt):
		fnName = gentypehelper.IdMapElasticToPlainSingle
	case f.TypeId.IsSingle(f.Opt) && f.TypeId.IsOptional(f.Opt):
		fnName = gentypehelper.IdMapElasticToPlainSingleOptional
	case !f.TypeId.IsSingle(f.Opt) && !f.TypeId.IsOptional(f.Opt):
		fnName = gentypehelper.IdMapElasticToPlainMultiple
	case !f.TypeId.IsSingle(f.Opt) && f.TypeId.IsOptional(f.Opt):
		fnName = gentypehelper.IdMapElasticToPlainMultipleOptional
	}

	return jen.Qual(gentypehelperQual, fnName)
}
