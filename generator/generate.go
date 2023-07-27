package generator

import (
	"encoding/hex"
	"strings"
	"unicode"
	"unicode/utf8"

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

func (g GeneratorOption) Gen(f *jen.File) {
	ctx := &generatorContext{
		generatorOption: g,
		file:            f,
	}
	ctx.Gen()
}

type generatorContext struct {
	generatorOption GeneratorOption
	file            *jen.File
	globalState     *globalState
	localState      localState
}

type globalState struct {
	generatedTypes map[string]mapping.Property // unused. future update may use this to dedup generated types.
}

type localState struct {
	fieldName []string
	prop      mapping.Property
	propOpt   PropertyOption
	dynamic   option.Option[mapping.DynamicMapping]
}

func (c *generatorContext) Gen() {
	newCtx := *c
	newCtx.globalState = &globalState{
		generatedTypes: make(map[string]mapping.Property),
	}
	newCtx.localState = localState{
		fieldName: append(c.localState.fieldName, c.generatorOption.RootTypeName),
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
			TypeName: pascalCase(escapeNonId(c.generatorOption.RootTypeName)),
			Children: c.generatorOption.MappingOption,
		},
		dynamic: c.generatorOption.Mapping.Dynamic.Option,
	}

	genObjectLike(&newCtx, false)
}

func (c *generatorContext) getTypeName() string {
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
func (c *generatorContext) next(
	fieldName string,
	prop mapping.Property,
	dynamic option.Option[mapping.DynamicMapping],
) *generatorContext {
	return &generatorContext{
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

func (c *generatorContext) IsOptional() bool {
	return c.localState.propOpt.IsOptional.Or(c.generatorOption.DefaultOption.IsOptional).OrElse(func() option.Option[bool] {
		return c.generatorOption.DefaultOption.PerTypDefault[mapping.GetTypeName(c.localState.prop)].IsOptional
	}).Value()
}
func (c *generatorContext) IsSingle() bool {
	return c.localState.propOpt.IsSingle.Or(c.generatorOption.DefaultOption.IsSingle).OrElse(func() option.Option[bool] {
		return c.generatorOption.DefaultOption.PerTypDefault[mapping.GetTypeName(c.localState.prop)].IsSingle
	}).Value()
}
func (c *generatorContext) PreferStringBoolean() bool {
	return c.localState.propOpt.PreferStringBoolean.Or(c.generatorOption.DefaultOption.PreferStringBoolean).Value()
}
func (c *generatorContext) PreferMarshalDateToNumber() bool {
	return c.localState.propOpt.PreferMarshalDateToNumber.Or(c.generatorOption.DefaultOption.PreferMarshalDateToNumber).Value()
}

var goOps = []string{
	">>=", "<<=", "...", "&^=", "||", "|=", "^=",
	">>", ">=", "==", "<=", "<<", "<-", ":=",
	"/=", "-=", "--", "+=", "++", "*=", "&^",
	"&=", "&&", "%=", "!=", "~", "}", "|",
	"{", "^", "]", "[", ">", "=", "<",
	";", ":", "/", ".", "-", ",", "+",
	"*", ")", "(", "&", "%", "!",
}

func escapeNonId(v string) string {
	builder := strings.Builder{}
	builder.Grow(len(v))
	var i int
LOOP:
	for i < len(v) {
		// As per the Go programming specification,
		// operators are listed as goOps.
		// https://go.dev/ref/spec#Operators_and_punctuation
		for _, op := range goOps {
			if strings.HasPrefix(v[i:], op) {
				builder.WriteByte('u')
				for _, letter := range []byte(op) {
					builder.WriteString("00")
					builder.WriteString(hex.EncodeToString([]byte{letter}))
				}
				i += len(op)
				continue LOOP
			}
		}
		r, size := utf8.DecodeRuneInString(v[i:])
		i += size
		// As per Go programming specification.
		// identifier = letter { letter | unicode_digit }.
		// https://go.dev/ref/spec#Identifiers
		if !(i == 0 && (unicode.IsDigit(r) || r == '_')) && (unicode.IsLetter(r) || r == '_' || unicode.IsDigit(r)) {
			builder.WriteRune(r)
		} else {
			builder.WriteByte('u')
			builder.WriteString(hex.EncodeToString([]byte(string(r))))
		}
	}

	return builder.String()
}

func pascalCase(snakeCase string) string {
	out := ""
	for _, part := range strings.Split(snakeCase, "_") {
		if len(part) == 0 {
			continue
		}
		out += strings.ToUpper(part[:1]) + part[1:]
	}
	return out
}
