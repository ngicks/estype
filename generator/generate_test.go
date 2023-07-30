package generator

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_escapeNonId(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []string{"✨", "👩", "🐱‍🐉", "__foo_bar"} {
		quoted := strings.ToLower(strconv.QuoteToASCII(tc))
		quoted = strings.ReplaceAll(quoted, "\\", "")
		quoted = quoted[1 : len(quoted)-1]

		escaped := escapeNonId(string(tc))

		assert.Equal(quoted, escaped)
		t.Logf("%s, %s\n", quoted, escaped)
	}
}

func Test_exportName(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range [][2]string{{"__foo_bar", "Foo_bar__"}} {
		exported := exportName(tc[0])

		assert.Equal(tc[1], exported)
		t.Logf("%s, %s\n", tc[1], exported)
	}
}

func Test_pascalCase(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range [][2]string{{"Foo_bar__", "FooBar__"}} {
		cased := pascalCase(tc[0])

		assert.Equal(tc[1], cased)
		t.Logf("%s, %s\n", tc[1], cased)
	}
}
