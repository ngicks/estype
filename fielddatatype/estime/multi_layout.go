package estime

import (
	"fmt"
	"sort"
	"time"
)

// MultiLayout is a set of time token layouts which can be used with time.Parse() or time.ParseInLocation.
// Layouts are internally sorted descendent order, longer is placed first.
type MultiLayout struct {
	layouts []string
}

// NewMultiLayout clones, sorts, dedups and validates input layouts.
//
// Sorting is stable: It sorts layouts by length in descending order,
// breaking ties by string comparison in descending order.
//
// Each one of layouts must have at least a single Go std's time layout token.
// Otherwise it returns an error.
func NewMultiLayout(layouts []string) (MultiLayout, error) {
	// clone first.
	cloned := make([]string, len(layouts))
	copy(cloned, layouts)

	// sort. longer is first.
	sort.Slice(cloned, func(i, j int) bool {
		iLen := len(cloned[i])
		jLen := len(cloned[j])
		if iLen != jLen {
			return iLen > jLen
		} else {
			return cloned[i] >= cloned[j]
		}
	})

	//dedup validating
	out := cloned[:0]
	for i := 0; i < len(cloned); i++ {
		if len(out) == 0 || cloned[i] != out[len(out)-1] {
			if !HasGoTimeToken(cloned[i]) {
				return MultiLayout{}, fmt.Errorf(
					"format %s does not contain any valid time layout tokens",
					cloned[i],
				)
			}
			out = append(out, cloned[i])
		}
	}

	return MultiLayout{
		layouts: out,
	}, nil
}

// NewMultiLayoutUnsafe returns MultiLayout without any safety checking
// which is otherwise employed in NewMultiLayout.
//
// Do not use this function unless a caller knows layouts are
// sorted, being valid that the time.Parse cause no error
// and dedupped.
func NewMultiLayoutUnsafe(layouts []string) MultiLayout {
	return MultiLayout{
		layouts: layouts,
	}
}

func (l MultiLayout) Clone() []string {
	out := make([]string, len(l.layouts))
	copy(out, l.layouts)
	return out
}

func (l MultiLayout) AddLayout(layout ...string) (MultiLayout, error) {
	return NewMultiLayout(append(layout, l.layouts...))
}

func (l MultiLayout) Parse(value string) (time.Time, error) {
	var lastError error
	for _, layout := range l.layouts {
		t, err := time.Parse(layout, value)
		if err == nil {
			return t, nil
		}
		lastError = err
	}
	return time.Time{}, lastError
}

func (l MultiLayout) ParseInLocation(value string, loc *time.Location) (time.Time, error) {
	var lastError error
	for _, layout := range l.layouts {
		t, err := time.ParseInLocation(layout, value, loc)
		if err == nil {
			return t, nil
		}
		lastError = err
	}
	return time.Time{}, lastError
}
