package codegentesthelper

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

const initialGoMod = `module testcodegen

go 1.20
`

func prepareTempDir(cfg Option) string {
	var tempDir = os.TempDir()
	if cfg.TempDir == "" {
		tempDir = cfg.TempDir
	}

	created, err := os.MkdirTemp(tempDir, "")
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepath.Join(created, "go.mod"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(f, bytes.NewBufferString(initialGoMod))
	if err != nil {
		panic(err)
	}

	return created
}
