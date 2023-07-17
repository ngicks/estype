package optionalstring

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type nodeTestCase struct {
	input    *treeNode
	expected []RawString
}

func TestNodeFlatten(t *testing.T) {
	cases := []nodeTestCase{
		{
			input: &treeNode{ // [A]
				typ: nonOptional,
				left: &treeNode{
					typ:   optional,
					value: []textNode{{typ: Normal, value: "A"}},
				},
			},
			expected: []RawString{
				{},
				{{typ: Normal, value: "A"}},
			},
		},
		{
			input: &treeNode{ // [A]B
				typ: nonOptional,
				left: &treeNode{
					typ:   optional,
					value: []textNode{{typ: Normal, value: "A"}},
				},
				right: &treeNode{
					typ:   nonOptional,
					value: []textNode{{typ: Normal, value: "B"}},
				},
			},
			expected: []RawString{
				{{typ: Normal, value: "B"}},
				{{typ: Normal, value: "A"}, {typ: Normal, value: "B"}},
			},
		},
		{
			input: &treeNode{ // A[B]C
				typ:   nonOptional,
				value: []textNode{{typ: Normal, value: "A"}},
				left: &treeNode{
					typ:   optional,
					value: []textNode{{typ: Normal, value: "B"}},
				},
				right: &treeNode{
					typ:   nonOptional,
					value: []textNode{{typ: Normal, value: "C"}},
				},
			},
			expected: []RawString{
				{{typ: Normal, value: "A"}, {typ: Normal, value: "C"}},
				{{typ: Normal, value: "A"}, {typ: Normal, value: "B"}, {typ: Normal, value: "C"}},
			},
		},
		{
			input: &treeNode{ // [A[B]C]
				typ: nonOptional,
				left: &treeNode{
					typ:   optional,
					value: []textNode{{typ: Normal, value: "A"}},
					left: &treeNode{
						typ:   optional,
						value: []textNode{{typ: Normal, value: "B"}},
					},
					right: &treeNode{
						typ:   nonOptional,
						value: []textNode{{typ: Normal, value: "C"}},
					},
				},
			},
			expected: []RawString{
				{},
				{{typ: Normal, value: "A"}, {typ: Normal, value: "C"}},
				{{typ: Normal, value: "A"}, {typ: Normal, value: "B"}, {typ: Normal, value: "C"}},
			},
		},
	}

	for _, tc := range cases {
		flattened := tc.input.Flatten()

		sort.Slice(tc.expected, func(i, j int) bool {
			return tc.expected[i].String() < tc.expected[j].String()
		})
		sort.Slice(flattened, func(i, j int) bool {
			return flattened[i].String() < flattened[j].String()
		})

		if len(tc.expected) != len(flattened) {
			t.Errorf("not same len. expected = %+#v, actual = %+#v", tc.expected, flattened)
		}
		for idx := range tc.expected {
			if diff := cmp.Diff(tc.expected[idx].String(), flattened[idx].String()); diff != "" {
				t.Errorf("not equal. diff = %s", diff)
			}
		}
	}
}
