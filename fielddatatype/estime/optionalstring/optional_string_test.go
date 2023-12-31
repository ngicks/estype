package optionalstring_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/ngicks/estype/fielddatatype/estime/optionalstring"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type variantsTestCases struct {
	input  string
	output []string
}

func TestMakeVariantsOptionalString(t *testing.T) {
	cases := []variantsTestCases{
		{
			input: `[YYYY[-M]M]-DDTHH:mm:ss.SSSZ`,
			output: []string{
				`YYYY-MM-DDTHH:mm:ss.SSSZ`,
				`YYYYM-DDTHH:mm:ss.SSSZ`,
				`-DDTHH:mm:ss.SSSZ`,
			},
		},
		{
			input: `YYYY-MM-DD'T'HH[:mm[:ss.SSS]]`,
			output: []string{
				`YYYY-MM-DD'T'HH:mm:ss.SSS`,
				`YYYY-MM-DD'T'HH:mm`,
				`YYYY-MM-DD'T'HH`,
			},
		},
		{
			input: `YYYY-MM-DD[THH[:mm[:ss.SSS]]][Z]`,
			output: []string{
				`YYYY-MM-DDTHH:mm:ss.SSSZ`,
				`YYYY-MM-DDTHH:mm:ss.SSS`,
				`YYYY-MM-DDTHH:mmZ`,
				`YYYY-MM-DDTHH:mm`,
				`YYYY-MM-DDTHHZ`,
				`YYYY-MM-DDTHH`,
				`YYYY-MM-DDZ`,
				`YYYY-MM-DD`,
			},
		},
		{
			input: `YYYY-MM-DD[THH[:mm[:ss.SSS]]]a[Z]`,
			output: []string{
				`YYYY-MM-DDTHH:mm:ss.SSSaZ`,
				`YYYY-MM-DDTHH:mm:ss.SSSa`,
				`YYYY-MM-DDTHH:mmaZ`,
				`YYYY-MM-DDTHH:mma`,
				`YYYY-MM-DDTHHaZ`,
				`YYYY-MM-DDTHHa`,
				`YYYY-MM-DDaZ`,
				`YYYY-MM-DDa`,
			},
		},
		{
			input: `[YYYY'[-M]'M]-DDTHH:mm:ss.SSSZ`,
			output: []string{
				`YYYY'[-M]'M-DDTHH:mm:ss.SSSZ`,
				`-DDTHH:mm:ss.SSSZ`,
			},
		},
		{
			input: `A[B]C`,
			output: []string{
				`ABC`,
				`AC`,
			},
		},
		{
			input: `ABC`,
			output: []string{
				`ABC`,
			},
		},
		{
			input: `A'B'C`,
			output: []string{
				`A'B'C`,
			},
		},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("case: %s", testCase.input), func(t *testing.T) {
			result, err := optionalstring.EnumerateOptionalString(testCase.input)
			require.NoError(t, err)
			sort.Strings(result)
			sort.Strings(testCase.output)
			assert.Equal(t, testCase.output, result)
		})
	}
}

func TestOptionalNonClosing(t *testing.T) {
	cases := []string{
		`foobar[baz[qux[`,
		`foobar\[baz[qux\[`,
		`foobarbaz]qux[]`,
		`foobarbaz\]qux[\]`,
		`foobarbaz\]qux\[]`,
		`foobar[ba[zq]ux`,
	}

	for _, input := range cases {
		_, err := optionalstring.EnumerateOptionalString(input)
		require.Error(t, err)
	}
}
