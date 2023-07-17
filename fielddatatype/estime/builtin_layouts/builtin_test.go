package builtinlayouts_test

import (
	"sort"
	"testing"
	"time"

	builtinlayouts "github.com/ngicks/estype/fielddatatype/estime/builtin_layouts"
	"github.com/stretchr/testify/assert"
)

func TestBuiltinFormats(t *testing.T) {
	assert := assert.New(t)

	keys := make([]string, len(builtinlayouts.BuiltinLayouts))
	for k := range builtinlayouts.BuiltinLayouts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	caseCount := 0
	failCount := 0

	for _, k := range keys {
		for _, layout := range builtinlayouts.BuiltinLayouts[k].Clone() {
			caseCount++
			didErr := false

			t.Logf("layout = %s", layout)

			for _, inputTime := range []time.Time{
				// input must not have any zero value for its input.
				// Do not place 1 for month and day, 0 for other fields.
				time.Date(2023, 7, 5, 3, 1, 8, 3, time.UTC),
				time.Date(2023, 12, 27, 18, 22, 53, 123456, time.UTC),
			} {

				serialized := inputTime.Format(layout)
				t.Logf("serialized to %s", serialized)
				decodedBack, err := time.Parse(layout, serialized)
				if !assert.NoError(err) {
					didErr = true
					continue
				}
				serializedAgain := decodedBack.Format(layout)
				decodedBackAgain, err := time.Parse(layout, serializedAgain)
				if !assert.NoError(err) {
					didErr = true
					continue
				}

				if !decodedBack.Equal(decodedBackAgain) {
					didErr = true
					t.Errorf(
						"not equal.\nexpected = %+#v\nactual   = %+#v\nserialized = %s\nlayout = %s, converted layout = %s",
						decodedBack, decodedBackAgain, serializedAgain, layout, layout,
					)
					continue
				}

				t.Logf(
					"input = %s, decoded to = %s",
					inputTime.Format(time.RFC3339Nano), decodedBackAgain.Format(time.RFC3339Nano),
				)
				year, month, day := decodedBackAgain.Date()
				var (
					hour    = decodedBackAgain.Hour()
					minute  = decodedBackAgain.Minute()
					sec     = decodedBackAgain.Second()
					nanoSec = decodedBackAgain.Nanosecond()
				)
				if year != 0 {
					assert.Equal(inputTime.Year(), year)
				}
				if month != time.January {
					assert.Equal(inputTime.Month(), month)
				}
				if day != 1 {
					assert.Equal(inputTime.Day(), day)
				}
				if hour != 0 {
					assert.Equal(inputTime.Hour(), hour)
				}
				if minute != 0 {
					assert.Equal(inputTime.Minute(), minute)
				}
				if sec != 0 {
					assert.Equal(inputTime.Second(), sec)
				}
				if nanoSec != 0 {
					assert.True(inputTime.Nanosecond()-nanoSec < 1000)
				}
			}
			if didErr {
				failCount++
			}
		}
	}

	t.Logf("tested %d cases and failed %d cases", caseCount, failCount)
}
