package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
	"github.com/ngicks/estype/generator"
	"github.com/ngicks/estype/helper/util"
	"github.com/ngicks/estype/spec/indices/indexstate"
)

var (
	outFile     = flag.String("o", "", "path to out file. defaults to stdout")
	configFile  = flag.String("c", "", "path to config. see definition of generate.GeneratorOption.")
	mappingFile = flag.String("m", "", "path to mapping.json.")
	packagePath = flag.String("p", "", "generated package path.")
)

func main() {
	flag.Parse()

	out := util.OpenOutMust(*outFile)
	defer out.Close()

	conf, err := os.Open(*configFile)
	if err != nil {
		panic(err)
	}
	defer conf.Close()

	mappingFile, err := os.Open(*mappingFile)
	if err != nil {
		panic(err)
	}
	defer mappingFile.Close()

	var generateOpt generator.GeneratorOption
	if err := json.NewDecoder(conf).Decode(&generateOpt); err != nil {
		panic(err)
	}

	var mappingIndex map[string]indexstate.IndexState
	if err := json.NewDecoder(mappingFile).Decode(&mappingIndex); err != nil {
		panic(err)
	}
	generateOpt.Mapping = mappingIndex[generateOpt.RootTypeName].Mappings.Value()

	generateOpt.GenerateTypeName = generator.ChainFieldName

	printOption(generateOpt)

	f := jen.NewFilePath(*packagePath)
	ctx := generateOpt.NewContext(f)
	ctx.Gen()

	if err := f.Render(out); err != nil {
		panic(err)
	}
}

func printOption(opt generator.GeneratorOption) {
	type printableConf struct {
		RootTypeName  string
		DefaultOption generator.DefaultOption
		MappingOption generator.MappingOption
	}
	printable := printableConf{
		RootTypeName:  opt.RootTypeName,
		DefaultOption: opt.DefaultOption,
		MappingOption: opt.MappingOption,
	}
	bin, err := json.MarshalIndent(printable, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bin))
}
