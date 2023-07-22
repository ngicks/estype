package test

//go:generate go run ../../cmd/genestype/gen.go -c ./testdata/all_opt.json -m ./testdata/all.json -o ./all.go -p github.com/ngicks/estype/generator/test
//go:generate go run ../../cmd/genestype/gen.go -c ./testdata/all_optional_opt.json -m ./testdata/all.json -o ./all_optional.go -p github.com/ngicks/estype/generator/test
//go:generate go run ../../cmd/genestype/gen.go -c ./testdata/dynamic_opt.json -m ./testdata/dynamic.json -o ./dynamic.go -p github.com/ngicks/estype/generator/test
//go:generate go run ../../cmd/genestype/gen.go -c ./testdata/conversion_opt.json -m ./testdata/conversion.json -o ./conversion.go -p github.com/ngicks/estype/generator/test
