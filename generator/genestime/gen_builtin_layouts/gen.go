package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/ngicks/estype/fielddatatype/estime"
)

var (
	outFile = flag.String("o", "", "")
)

// This JSON is generated from text extracted from the Elasticsearch document web page.
// To get the JSON, visit the page: https://www.elastic.co/guide/en/elasticsearch/reference/8.8/mapping-date-format.html
// and execute this one-liner script in the console:
//
//	JSON.stringify(Object.fromEntries([...document.getElementById("built-in-date-formats").parentElement.parentElement.parentElement.parentElement.parentElement.querySelectorAll("dt .literal").values()].map(ele => [ele.innerText,[...ele.parentElement.parentElement.nextSibling.nextSibling.querySelectorAll(".literal").values()].map(ele => ele.innerText)])))
//
// Its page layout may change over time. Adjust it if so.
//
//go:embed builtin_formats.jsonc
var builtInFormats string

var commentLine = regexp.MustCompile(`^\s*//`)

func main() {
	flag.Parse()

	lines := strings.Split(builtInFormats, "\n")
	commentLineRemoved := lines[:0]
	// removing comment lines.
	for i := 0; i < len(lines); i++ {
		if !commentLine.MatchString(lines[i]) {
			if pos := strings.Index(lines[i], "//"); pos >= 0 {
				lines[i] = lines[i][:pos]
			}
			commentLineRemoved = append(commentLineRemoved, lines[i])
		}
	}
	lines = commentLineRemoved

	// Was jsonc. Now it is json since we removed comments.
	pureJson := strings.Join(lines, "\n")

	builtinFormatTable := make(map[string][]string)

	err := json.Unmarshal([]byte(pureJson), &builtinFormatTable)
	if err != nil {
		panic(err)
	}

	// making definition-ordered array
	bin := []byte(pureJson)
	skipStart := 0
	for i := 0; i < len(bin); i++ {
		if bin[i] == ':' && skipStart == 0 {
			skipStart = i
		}
		if bin[i] == ']' {
			copy(bin[skipStart:], bin[i+1:])
			bin = bin[:len(bin)-(i+1-skipStart)]
			i = skipStart
			skipStart = 0
		}
	}
	bin[0] = '['
	bin[len(bin)-2] = ']'

	formatOrder := make([]string, 0)
	err = json.Unmarshal(bin, &formatOrder)
	if err != nil {
		panic(err)
	}

	for _, v := range builtinFormatTable {
		for i := 0; i < len(v); i++ {
			v[i], _ = estime.ConvertTimeToken(v[i])
		}
	}

	var out io.Writer
	if outFile == nil || *outFile == "" {
		out = os.Stdout
	} else {
		f, err := os.Create(*outFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		out = f
	}

	f := jen.NewFilePath("github.com/ngicks/estype/fielddatatype/estime/builtin_layouts")

	f.HeaderComment("// Code generated by github.com/ngicks/estype/generator/genestime/gen_builtin_layouts/gen.go. DO NOT EDIT.")

	f.PackageComment("// builtinlayouts is collection of date formats built in Elasticsearch")
	f.PackageComment("")
	f.PackageComment("//go:generate go run ../../../generator/genestime/gen_builtin_layouts/gen.go -o layout.go")
	f.
		Var().
		Id("BuiltinLayouts").
		Op("=").
		Map(jen.String()).
		Qual("github.com/ngicks/estype/fielddatatype/estime", "MultiLayout").
		ValuesFunc(func(g *jen.Group) {
			for _, formatName := range formatOrder {
				g.Add(
					jen.
						Line().
						Lit(formatName).
						Op(":").
						Id("must").
						Call(
							jen.
								Qual("github.com/ngicks/estype/fielddatatype/estime", "NewMultiLayout").
								Call(
									jen.Index().String().ValuesFunc(func(g *jen.Group) {
										for _, format := range builtinFormatTable[formatName] {
											g.Add(jen.Line().Lit(format))
										}
										g.Line()
									}),
								),
						),
				)
			}
		})

	if err := f.Render(out); err != nil {
		panic(err)
	}
}