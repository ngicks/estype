package estime

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestEsTime(t *testing.T) {
	assert := assert.New(t)

	p, err := FromJavaDateTimeLike(
		[]string{"yyyy-MM-dd['T'HH:mm:ss.SSSZ]", "yy-M-d['T'HH:mm:ss.SSSZ]"},
		Millis,
	)
	assert.NoError(err)

	diff := cmp.Diff(
		[]string{
			"2006-01-02T15:04:05.000Z0700",
			"06-1-2T15:04:05.000Z0700",
			"2006-01-02",
			"06-1-2",
		},
		p.multiLayout.layouts,
	)
	assert.Empty(diff)
}
