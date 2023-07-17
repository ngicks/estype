package estime_test

import (
	"sort"
	"testing"

	"github.com/ngicks/estype/fielddatatype/estime"
	"github.com/stretchr/testify/assert"
)

func TestHasGoTimeToken(t *testing.T) {
	assert := assert.New(t)

	timeTokens := make([]string, len(goTimeToken))
	var i = 0
	for k := range goTimeToken {
		timeTokens[i] = k
		i++
	}
	sort.Strings(timeTokens)
	// Make log or error messages stable.

	for _, token := range timeTokens {
		assert.True(estime.HasGoTimeToken("____"+token), "____"+token)
		assert.True(estime.HasGoTimeToken("___"+token+"__a___aaa"), "___"+token+"__a___aaa")
		if len(token) >= 2 {
			former, latter := token[:1], token[1:]
			if _, ok := goTimeToken[former]; ok {
				continue
			}
			if _, ok := goTimeToken[latter]; ok {
				continue
			}
			assert.False(estime.HasGoTimeToken(former + "_" + latter))
		}
	}
}

var goTimeToken = map[string]struct{}{
	"January":   {},
	"Jan":       {},
	"1":         {},
	"01":        {},
	"Monday":    {},
	"Mon":       {},
	"2":         {},
	"_2":        {},
	"02":        {},
	"__2":       {},
	"002":       {},
	"15":        {},
	"3":         {},
	"03":        {},
	"4":         {},
	"04":        {},
	"5":         {},
	"05":        {},
	"2006":      {},
	"06":        {},
	"PM":        {},
	"pm":        {},
	"MST":       {},
	"Z0700":     {}, // prints Z for UTC
	"Z070000":   {},
	"Z07":       {},
	"Z07:00":    {}, // prints Z for UTC
	"Z07:00:00": {},
	"-0700":     {}, // always numeric
	"-070000":   {},
	"-07":       {}, // always numeric
	"-07:00":    {}, // always numeric
	"-07:00:00": {},
	".0":        {},
	".9":        {},
}
