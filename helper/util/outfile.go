package util

import "os"

func OpenOutMust(outFile string) *os.File {
	if f, err := OpenOut(outFile); err != nil {
		panic(err)
	} else {
		return f
	}
}

func OpenOut(outFile string) (*os.File, error) {
	if outFile == "" || outFile == "-" || outFile == "--" {
		return os.Stdout, nil
	} else {
		f, err := os.Create(outFile)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
}
