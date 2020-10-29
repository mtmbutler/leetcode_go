package maxdepth

import "testing"

type Case struct {
	s        string
	expected int
}

var cases = []Case{
	{"(1+(2*3)+((8)/4))+1", 3},
	{"(1)+((2))+(((3)))", 3},
	{"1+(2*3)/(2-1)", 1},
	{"1", 0},
	{"", 0},
	{"()", 1},
	{"(())", 2},
}

func TestMaxDepth(t *testing.T) {
	for _, c := range cases {
		result := maxDepth(c.s)
		if result != c.expected {
			t.Errorf(
				"maxDepth(%v) == %v, want %v",
				c.s, result, c.expected,
			)
		}
	}
}
