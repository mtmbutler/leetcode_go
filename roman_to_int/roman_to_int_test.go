package romantoint

import "testing"

type Case struct {
	s        string
	expected int
}

var cases = []Case{
	{"I", 1},
	{"", 0},
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"CMXLIX", 949},
}

func TestRomanToInt(t *testing.T) {
	for _, c := range cases {
		got := romanToInt(c.s)
		if got != c.expected {
			t.Errorf(
				"romanToInt(%s) == %d, want %d",
				c.s,
				got,
				c.expected,
			)
		}
	}
}
