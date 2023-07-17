package estime_test

import (
	"testing"
	"time"

	"github.com/ngicks/estype/fielddatatype/estime"
	"github.com/stretchr/testify/assert"
)

var PDT = time.FixedZone("PDT", -7*60*60)
var sampleInput = time.Date(2001, 7, 4, 12, 8, 56, 235000000, PDT)

func TestFormat(t *testing.T) {
	assert := assert.New(t)

	type testTy struct {
		layout    string
		converted string
		output    string
	}

	for _, tc := range []testTy{
		{
			"yyyy.MM.dd 'at' HH:mm:ss z",
			"2006.01.02 at 15:04:05 MST",
			"2001.07.04 at 12:08:56 PDT",
		},
		{
			"EEE, MMM d, ''yy",
			"Mon, Jan 2, '06",
			"Wed, Jul 4, '01",
		},
		{
			"EEEE, MMMM, DDD",
			"Monday, January, 002",
			"Wednesday, July, 185",
		},
		{
			"h:mm a",
			"3:04 PM",
			"12:08 PM",
		},
		{
			"EEE, d MMM yyyy HH:mm:ss Z",
			"Mon, 2 Jan 2006 15:04:05 Z0700",
			"Wed, 4 Jul 2001 12:08:56 -0700",
		},
		{
			"yyMMddHHmmssZ",
			"060102150405Z0700",
			"010704120856-0700",
		},
		{
			"yyyy-MM-dd'T'HH:mm:ss.SSSZ",
			"2006-01-02T15:04:05.000Z0700",
			"2001-07-04T12:08:56.235-0700",
		},
		{
			"yyyy-MM-dd'T'HH:mm:ss.SSSXXX",
			"2006-01-02T15:04:05.000Z070000",
			"2001-07-04T12:08:56.235-070000",
		},
	} {
		layout, err := estime.ConvertTimeToken(tc.layout)
		assert.NoError(err, "input = %s", tc.layout)
		assert.Equal(tc.converted, layout)
		out := sampleInput.Format(layout)
		assert.Equal(tc.output, out)
	}
}
