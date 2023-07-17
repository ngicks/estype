package fielddatatype_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ngicks/estype/fielddatatype"
	"github.com/stretchr/testify/require"
)

type TestBool struct {
	A    *fielddatatype.Boolean
	B    *fielddatatype.Boolean
	Astr *fielddatatype.BooleanStr
	Bstr *fielddatatype.BooleanStr
}

func TestBoolean(t *testing.T) {
	for _, testCase := range [][4]bool{
		{true, true, true, true},
		{true, false, true, false},
		{false, true, false, true},
		{false, false, false, false},
	} {
		testBool := TestBool{}
		err := json.Unmarshal(
			[]byte(fmt.Sprintf(
				`{"A": %t, "B": "%t", "Astr": %t, "Bstr": "%t"}`,
				testCase[0],
				testCase[1],
				testCase[2],
				testCase[3],
			)),
			&testBool,
		)
		require.NoError(t, err)
		require.Equal(
			t,
			TestBool{
				A:    escape(fielddatatype.Boolean(testCase[0])),
				B:    escape(fielddatatype.Boolean(testCase[1])),
				Astr: escape(fielddatatype.BooleanStr(testCase[2])),
				Bstr: escape(fielddatatype.BooleanStr(testCase[3])),
			},
			testBool,
		)
		bin, _ := json.Marshal(testBool)

		testBool2 := TestBool{}
		err = json.Unmarshal(bin, &testBool2)
		require.NoError(t, err)
		require.Equal(t, testBool, testBool2)
	}
}

func escape[T any](t T) *T {
	return &t
}

func TestBooleanEmptyString(t *testing.T) {
	type testUnmarshal struct {
		Lit *fielddatatype.Boolean
		Str *fielddatatype.BooleanStr
	}
	testBool := testUnmarshal{
		Lit: escape(fielddatatype.Boolean(true)),
		Str: escape(fielddatatype.BooleanStr(true)),
	}
	err := json.Unmarshal([]byte(`{"Lit": "",  "Str": ""}`), &testBool)
	require.NoError(t, err)

	require.Equal(t, fielddatatype.Boolean(false), *testBool.Lit)
	require.Equal(t, fielddatatype.BooleanStr(false), *testBool.Str)
}

func TestBooleanInvalidInput(t *testing.T) {
	type TestBool struct {
		A fielddatatype.Boolean
	}
	type TestBoolStr struct {
		A fielddatatype.BooleanStr
	}
	testBool := TestBool{}
	testBoolStr := TestBoolStr{}

	testBooleanUnmarshal(t, testBool)
	testBooleanUnmarshal(t, testBoolStr)

	var err error
	var invalidTypeErr *fielddatatype.InvalidTypeError
	err = testBool.A.UnmarshalJSON([]byte(`dawju9813`))
	require.ErrorAs(t, err, &invalidTypeErr)
}

func testBooleanUnmarshal[T any](t *testing.T, testBool T) {
	var invalidTypeError *fielddatatype.InvalidTypeError
	for _, testCase := range []string{
		`{"A": "foo"}`,
		`{"A": 123}`,
		`{"A": 123.5}`,
	} {
		err := json.Unmarshal([]byte(testCase), &testBool)
		require.ErrorAs(t, err, &invalidTypeError)
		// checking just that it does not panic
		require.Condition(t, func() bool { return strings.HasPrefix(err.Error(), "invalid") })
	}
}

func TestBooleanString(t *testing.T) {
	esBoolean := fielddatatype.Boolean(false)
	esBooleanStr := fielddatatype.BooleanStr(false)

	require.Equal(t, "false", esBoolean.String())
	require.Equal(t, "false", esBooleanStr.String())

	esBoolean = fielddatatype.Boolean(true)
	esBooleanStr = fielddatatype.BooleanStr(true)

	require.Equal(t, "true", esBoolean.String())
	require.Equal(t, "true", esBooleanStr.String())
}
