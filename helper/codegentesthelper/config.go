package codegentesthelper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

type Option struct {
	TempDir string
}

func ParseEnv() Option {
	return Option{
		TempDir: os.Getenv("TEMP_DIR"),
	}
}

func (o Option) ToCtx() Ctx {
	return Ctx{
		Option: o,
	}
}

type Ctx struct {
	Option Option
	Dir    string
}

func (c *Ctx) PrepareDir() {
	c.Dir = prepareTempDir(c.Option)
}

func (c *Ctx) WriteMain(fn func(w io.Writer) error) {
	if c.Dir == "" {
		c.PrepareDir()
	}

	f, err := os.Create(filepath.Join(c.Dir, "main.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := fn(f); err != nil {
		panic(err)
	}

	c.goModTidy()
}

func (c *Ctx) Gen(fn func(w io.Writer) error) {
	if c.Dir == "" {
		c.PrepareDir()
	}
	f, err := os.Create(filepath.Join(c.Dir, "gen.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := fn(f); err != nil {
		panic(err)
	}

	c.goModTidy()
}

func (c *Ctx) goModTidy() {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = c.Dir
	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("%w: go mod tidy failed", err))
	}
}

func (c *Ctx) Run(ctx context.Context) (combinedOutput []byte, err error) {
	if c.Dir == "" {
		c.PrepareDir()
	}
	outBuf := new(bytes.Buffer)
	// We are not performing `go run .` because we'll remove all generated files including built artifacts.
	cmd := exec.CommandContext(ctx, "go", "build", "-o", "main", "./main.go")
	cmd.Dir = c.Dir
	cmd.Stdout = outBuf
	cmd.Stderr = outBuf
	err = cmd.Run()
	if err != nil {
		return outBuf.Bytes(), err
	}

	outBuf.Reset()

	cmd = exec.CommandContext(ctx, "./main")
	cmd.Dir = c.Dir
	cmd.Stdout = outBuf
	cmd.Stderr = outBuf
	err = cmd.Run()
	return outBuf.Bytes(), err
}

func (c *Ctx) RemoveAll() error {
	return os.RemoveAll(c.Dir)
}
