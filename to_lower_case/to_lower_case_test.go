package tolowercase

import "testing"

type Case struct {
	str      string
	expected string
}

var cases = []Case{
	{"", ""},
	{".", "."},
	{"Hello", "hello"},
	{"here", "here"},
	{"LOVELY", "lovely"},
}

func TestToLowerCase(t *testing.T) {
	for _, c := range cases {
		got := toLowerCase(c.str)
		if got != c.expected {
			t.Errorf(
				"toLowerCase(%s) == %s, want %s",
				c.str,
				got,
				c.expected,
			)
		}
	}
}
