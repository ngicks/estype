package generator

import (
	"strings"

	"github.com/ngicks/estype/fielddatatype/estime"
	builtinlayouts "github.com/ngicks/estype/fielddatatype/estime/builtin_layouts"
	"github.com/ngicks/estype/generator/genestime"
	"github.com/ngicks/estype/spec/mapping"
)

const (
	builtinEsDateQual = "github.com/ngicks/estype/fielddatatype/estime/builtin"
)

func genDate(ctx *generatorContext, dryRun bool) typeId {
	var formats string
	isNano := false
	switch x := ctx.localState.prop.Val.(type) {
	case mapping.DateProperty:
		if x.Format.IsDefined() {
			formats = x.Format.Value()
		}
	case mapping.DateNanosProperty:
		isNano = true
		if x.Format.IsDefined() {
			formats = x.Format.Value()
		}
	case mapping.DateRangeProperty:
		if x.Format.IsDefined() {
			formats = x.Format.Value()
		}
	}

	if len(formats) == 0 {
		var id string
		if ctx.generatorOption.DefaultOption.PreferMarshalDateToNumber.Value() {
			id = "DefaultNum"
		} else {
			id = "Default"
		}
		if isNano {
			id += "Nano"
		}
		return typeId{
			Qualifier: builtinEsDateQual,
			Id:        id,
		}
	}

	stringFormats, numFormat := parseFormats(formats)
	tyId, ok := pregeneratedDate(stringFormats, numFormat)
	if ok {
		return tyId
	}

	tyId = typeId{
		Id: ctx.getTypeName(),
	}
	// This is not great that FromJavaDateTimeLike is used to just parse + validation.
	layouts, err := estime.FromJavaDateTimeLike(stringFormats, "")
	if err != nil {
		panic(err)
	}

	if !dryRun {
		numParser := estime.NumParser(numFormat)
		genestime.GeneratorDef{
			TyName:          tyId.Id,
			MultiLayout:     estime.NewMultiLayoutUnsafe(layouts.Layout()),
			NumParser:       numParser,
			MarshalToNumber: ctx.PreferMarshalDateToNumber(),
		}.Gen(ctx.file)
	}
	return tyId
}

// parseFormats parses double-vertical-line (`||`) separated formats into stringFormats and number formats,
// converting builtin format into Go time layout.
func parseFormats(formats string) (stringFormat []string, numFormat string) {
	if len(formats) == 0 {
		return
	}
	for _, format := range strings.Split(formats, "||") {
		switch format {
		case builtinlayouts.EpochMillis, builtinlayouts.EpochSecond:
			numFormat = format
		default:
			layouts, ok := builtinlayouts.BuiltinLayouts[format]
			if ok {
				stringFormat = append(stringFormat, layouts.Clone()...)
			} else {
				stringFormat = append(stringFormat, format)
			}
		}
	}
	return
}

func pregeneratedDate(stringFormat []string, numFormat string) (typeId, bool) {
	if len(numFormat) > 0 && len(stringFormat) == 0 {
		return typeId{Qualifier: builtinEsDateQual, Id: pascalCase(numFormat)}, true
	}
	if numFormat == "" && len(stringFormat) == 1 {
		_, ok := builtinlayouts.BuiltinLayouts[stringFormat[0]]
		if ok {
			return typeId{Qualifier: builtinEsDateQual, Id: pascalCase(stringFormat[0])}, true
		}
	}
	return typeId{}, false
}
