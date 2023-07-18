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
						stmt := jen.
							Id(field.Name).
							Op(":").
							Qual(elasticTypeQual, fromFuncName(plainFields[idx]))

							// at this point;
							// foo: elastic.From

						var input *jen.Statement
						if !plainFields[idx].IsObjectLike {
							input = jen.Id("d").Dot(field.Name)
						} else {
							if plainFields[idx].Opt.IsSingle() {
								input = jen.Id("d").Dot(field.Name).Dot("ToRaw").Call()
							} else {
								fieldType := jen.
									Id(field.TypeId.TypeParam[0].Id)
								var mapperId string
								if plainFields[idx].Opt.IsOptional() {
									mapperId = mapToRawPointerId
								} else {
									mapperId = mapToRawId
								}
								input = jen.Id(mapperId).Index(fieldType).Call(jen.Id("d").Dot(field.Name))
							}
						}

						if escaper := escaperId(plainFields[idx]); escaper != "" {
							input = jen.Id(escaper).Call(input)
						}

						stmt = stmt.Call(input)

						g.Add(stmt)
					}
					if !strict {
						g.Add(jen.Id(additionalPropId).Op(":").Id("d").Dot(additionalPropId))
					}
				},
			),
		),
	)
}

func fromFuncName(f structField) string {
	fromFnName := "From"
	if f.Opt.IsSingle() {
		fromFnName += "Single"
	} else {
		fromFnName += "Multiple"
	}
	if f.Opt.IsOptional() {
		fromFnName += "Pointer"
	}
	return fromFnName
}

func escaperId(f structField) string {
	if f.Opt.IsOptional() {
		if f.Opt.IsSingle() {
			return escapeValueId
		} else {
			return escapeSliceId
		}
	}
	return ""
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
					for _, field := range plainFields {
						stmt := jen.
							Id(field.Name).
							Op(":")

						var value *jen.Statement
						if !field.IsObjectLike {
							value = jen.Id("d").Dot(field.Name).Dot(toFuncName(field)).Call()
						} else {
							if field.Opt.IsSingle() {
								value = jen.Id("d").Dot(field.Name).Dot(toFuncName(field)).Call().Dot("ToPlain").Call()
							} else {
								fieldType := jen.Id(field.TypeId.Id)
								value = jen.Id(mapToPlainId).Index(fieldType).Call(jen.Id("d").Dot(field.Name).Dot("ValueMultiple").Call())
							}
						}

						if escaper := escaperId(field); escaper != "" {
							value = jen.Id(escaper).Call(value)
						}

						g.Add(stmt.Add(value))
					}
					if !strict {
						g.Add(jen.Id(additionalPropId).Op(":").Id("d").Dot(additionalPropId))
					}
				},
			),
		),
	)
}

func toFuncName(f structField) string {
	var toFuncName string
	if f.Opt.IsOptional() {
		toFuncName += "Plain"
	} else {
		toFuncName += "Value"
	}
	if f.Opt.IsSingle() {
		toFuncName += "Single"
	} else {
		toFuncName += "Multiple"
	}
	return toFuncName
}
