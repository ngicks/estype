package estime_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/estype/fielddatatype/estime"
	"github.com/stretchr/testify/assert"
)

func TestMultiLayout(t *testing.T) {
	assert := assert.New(t)

	ml, err := estime.NewMultiLayout([]string{
		"2006-01-02T15:04:05.000Z0700",
		"06-1-2T15:04:05.000Z0700",
		"2006-01-02",
		"06-1-2",
	})
	assert.NoError(err)

	assert.Empty(
		cmp.Diff(
			[]string{
				"2006-01-02T15:04:05.000Z0700",
				"06-1-2T15:04:05.000Z0700",
				"2006-01-02",
				"06-1-2",
			},
			ml.Clone(),
		),
	)

	ml, err = ml.AddLayout("Mon, 2 Jan 2006 15:04:05 Z0700")
	assert.NoError(err)

	assert.Empty(
		cmp.Diff(
			[]string{
				"Mon, 2 Jan 2006 15:04:05 Z0700",
				"2006-01-02T15:04:05.000Z0700",
				"06-1-2T15:04:05.000Z0700",
				"2006-01-02",
				"06-1-2",
			},
			ml.Clone(),
		),
	)

	tt, err := ml.Parse("Wed, 4 Jul 2001 12:08:56 -0700")
	assert.NoError(err)
	// The format looses milli secs and lower precision.
	assert.True(tt.Equal(sampleInput.Add(-235000000)), "diff = %s", cmp.Diff(tt, sampleInput))

	tt, err = ml.Parse("2001-07-04T12:08:56.235-0700")
	assert.NoError(err)
	assert.True(tt.Equal(sampleInput), "diff = %s", cmp.Diff(tt, sampleInput))

	tt, err = ml.Parse("01-7-4")
	assert.NoError(err)
	expected := time.Date(2001, 7, 4, 0, 0, 0, 0, time.UTC)
	assert.True(tt.Equal(expected), "diff = %s", cmp.Diff(tt, expected))

	tt, err = ml.ParseInLocation("01-7-4", PDT)
	assert.NoError(err)
	expected = time.Date(2001, 7, 4, 0, 0, 0, 0, PDT)
	assert.True(tt.Equal(expected), "diff = %s", cmp.Diff(tt, expected))
}

func TestMultiLayout_dedup(t *testing.T) {
	assert := assert.New(t)

	ml, err := estime.NewMultiLayout([]string{
		"2006-01-02T15:04:05.000Z0700",
		"2006-01-02T15:04:05.000Z0700",
		"2006-01-02",
		"2006-01-02",
	})
	assert.NoError(err)

	assert.Empty(
		cmp.Diff(
			[]string{
				"2006-01-02T15:04:05.000Z0700",
				"2006-01-02",
			},
			ml.Clone(),
		),
	)
}

func TestMultiLayout_error_for_non_time(t *testing.T) {
	assert := assert.New(t)

	_, err := estime.NewMultiLayout([]string{
		"yyyy.MM.dd 'at' HH:mm:ss z",
	})
	assert.Error(err)
}

func TestMultiLayout_sort(t *testing.T) {
	assert := assert.New(t)

	ml, err := estime.NewMultiLayout([]string{
		"2",
		"1",
		"02",
	})
	assert.NoError(err)
	assert.Empty(cmp.Diff([]string{"02", "2", "1"}, ml.Clone()))
}
