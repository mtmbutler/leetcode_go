package lenlastword

import "testing"

type Case struct {
	s        string
	expected int
}

var cases = []Case{
	{"Hello World", 5},
	{"foo bar baz", 3},
	{"Nospaceshere", 12},
	{"", 0},
	{"a", 1},
	{"a ", 1},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, c := range cases {
		got := lengthOfLastWord(c.s)
		if got != c.expected {
			t.Errorf(
				"lengthOfLastWord(%s) == %d, want %d",
				c.s,
				got,
				c.expected,
			)
		}
	}
}
