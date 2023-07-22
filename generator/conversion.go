package generator

import "github.com/dave/jennifer/jen"

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
		if plain.TypeId.IsOptional(plain.Opt) {
			return jen.Qual(jsonfieldTypeQual, "FromPointer").Call(fieldNameId)
		} else {
			return jen.Qual(jsonfieldTypeQual, "Defined").Call(fieldNameId)
		}
	case undefinedableTypeQual:
		if plain.TypeId.IsOptional(plain.Opt) {
			return jen.Qual(undefinedableTypeQual, "FromPointer").Call(fieldNameId)
		} else {
			return jen.Qual(undefinedableTypeQual, "Defined").Call(fieldNameId)
		}
	case elasticTypeQual:
		var input *jen.Statement
		if !plain.IsObjectLike {
			input = jen.Id("d").Dot(raw.Name)
		} else {
			if plain.Opt.IsSingle() {
				input = jen.Id("d").Dot(raw.Name).Dot("ToRaw").Call()
			} else {
				fieldType := jen.
					Id(raw.TypeId.TypeParam[0].Id)
				var mapperId string
				if plain.Opt.IsOptional() {
					mapperId = mapToRawPointerId
				} else {
					mapperId = mapToRawId
				}
				input = jen.Id(mapperId).Index(fieldType).Call(jen.Id("d").Dot(raw.Name))
			}
		}
		if escaper := escaperId(plain); escaper != "" {
			input = jen.Id(escaper).Call(input)
		}

		return jen.Qual(elasticTypeQual, fromFuncName(plain)).Call(input)
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
		var value *jen.Statement
		if !raw.IsObjectLike {
			value = jen.Id("d").Dot(raw.Name).Dot(toFuncName(raw)).Call()
		} else {
			if raw.TypeId.IsSingle(raw.Opt) {
				value = jen.Id("d").Dot(raw.Name).Dot(toFuncName(raw)).Call().Dot("ToPlain").Call()
			} else {
				fieldType := jen.Id(raw.TypeId.Id)
				value = jen.Id(mapToPlainId).Index(fieldType).Call(fieldNameId.Dot("ValueMultiple").Call())
			}
		}

		if escaper := escaperId(raw); escaper != "" {
			value = jen.Id(escaper).Call(value)
		}

		return value
	}
	panic("unknown")
}

func fromFuncName(f structField) string {
	fromFnName := "From"
	if f.TypeId.IsSingle(f.Opt) {
		fromFnName += "Single"
	} else {
		fromFnName += "Multiple"
	}
	if f.TypeId.IsOptional(f.Opt) {
		fromFnName += "Pointer"
	}
	return fromFnName
}

func escaperId(f structField) string {
	if f.TypeId.IsOptional(f.Opt) {
		if f.TypeId.IsSingle(f.Opt) {
			return escapeValueId
		} else {
			return escapeSliceId
		}
	}
	return ""
}

func toFuncName(f structField) string {
	var toFuncName string
	if f.TypeId.IsOptional(f.Opt) {
		toFuncName += "Plain"
	} else {
		toFuncName += "Value"
	}
	if f.TypeId.IsSingle(f.Opt) {
		toFuncName += "Single"
	} else {
		toFuncName += "Multiple"
	}
	return toFuncName
}
