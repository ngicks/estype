package generator

import (
	"strings"

	"github.com/dave/jennifer/jen"

	"github.com/ngicks/estype/spec/mapping"
	"github.com/ngicks/und/option"
)

type GeneratorOption struct {
	RootTypeName     string
	Mapping          mapping.TypeMapping
	GenerateTypeName GenerateTypeName
	DefaultOption    DefaultOption
	MappingOption    MappingOption
}

func (g GeneratorOption) NewContext(f *jen.File) *GeneratorContext {
	return &GeneratorContext{
		generatorOption: g,
		file:            f,
	}
}

type GeneratorContext struct {
	generatorOption GeneratorOption
	file            *jen.File
	globalState     *globalState
	localState      localState
}

type globalState struct {
	generatedTypes           map[string]mapping.Property // unused. future update may use this to dedup generated types.
	bufPoolGenerated         bool
	escapeValueGenerated     bool
	escapeSliceGenerated     bool
	mapToPlainGenerated      bool
	mapToRawGenerated        bool
	mapToRawPointerGenerated bool
}

type localState struct {
	fieldName []string
	prop      mapping.Property
	propOpt   PropertyOption
	dynamic   option.Option[mapping.DynamicMapping]
}

func (c *GeneratorContext) Gen() {
	newCtx := *c
	newCtx.globalState = &globalState{
		generatedTypes: make(map[string]mapping.Property),
	}
	newCtx.localState = localState{
		fieldName: append(c.localState.fieldName, pascalCase(c.generatorOption.RootTypeName)),
		prop: mapping.Property{
			Val: mapping.ObjectProperty{
				CorePropertyBase: mapping.CorePropertyBase{
					PropertyBase: mapping.PropertyBase{
						Properties: c.generatorOption.Mapping.Properties,
					},
				},
			},
		},
		propOpt: PropertyOption{
			TypeName: pascalCase(c.generatorOption.RootTypeName),
			Children: c.generatorOption.MappingOption,
		},
		dynamic: c.generatorOption.Mapping.Dynamic.Option,
	}

	genObjectLike(&newCtx, false)
}

func (c *GeneratorContext) getTypeName() string {
	return c.localState.propOpt.GetTypeName(
		func() string {
			return c.generatorOption.GenerateTypeName(
				c.localState.fieldName,
				mapping.GetTypeName(c.localState.prop),
			)
		},
	)
}

// next proceeds ctx one step deeper into properties of object like mapping.
func (c *GeneratorContext) next(
	fieldName string,
	prop mapping.Property,
	dynamic option.Option[mapping.DynamicMapping],
) *GeneratorContext {
	return &GeneratorContext{
		generatorOption: c.generatorOption,
		file:            c.file,
		globalState:     c.globalState,
		localState: localState{
			fieldName: append(c.localState.fieldName, fieldName),
			prop:      prop,
			propOpt:   c.localState.propOpt.Children[fieldName],
			dynamic:   dynamic.Or(c.localState.dynamic),
		},
	}
}

func (c *GeneratorContext) IsOptional() bool {
	return c.localState.propOpt.IsOptional.Or(c.generatorOption.DefaultOption.IsOptional).OrElse(func() option.Option[bool] {
		return c.generatorOption.DefaultOption.PerTypDefault[mapping.GetTypeName(c.localState.prop)].IsOptional
	}).Value()
}
func (c *GeneratorContext) IsSingle() bool {
	return c.localState.propOpt.IsSingle.Or(c.generatorOption.DefaultOption.IsSingle).OrElse(func() option.Option[bool] {
		return c.generatorOption.DefaultOption.PerTypDefault[mapping.GetTypeName(c.localState.prop)].IsSingle
	}).Value()
}
func (c *GeneratorContext) PreferStringBoolean() bool {
	return c.localState.propOpt.PreferStringBoolean.Or(c.generatorOption.DefaultOption.PreferStringBoolean).Value()
}
func (c *GeneratorContext) PreferMarshalDateToNumber() bool {
	return c.localState.propOpt.PreferMarshalDateToNumber.Or(c.generatorOption.DefaultOption.PreferMarshalDateToNumber).Value()
}

func pascalCase(snakeCase string) string {
	out := ""
	for _, part := range strings.Split(snakeCase, "_") {
		out += strings.ToUpper(part[:1]) + part[1:]
	}
	return out
}
