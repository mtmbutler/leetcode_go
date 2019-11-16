package zigzagconversion

import "testing"

type Case struct {
	s        string
	numRows  int
	expected string
}

var cases = []Case{
	{"", 2, ""},
	{"foobar", 6, "foobar"},
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	{"AB", 1, "AB"},
}

func TestConvert(t *testing.T) {
	for _, c := range cases {
		got := convert(c.s, c.numRows)
		if got != c.expected {
			t.Errorf(
				"convert(%s, %d) == %s, want %s",
				c.s,
				c.numRows,
				got,
				c.expected,
			)
		}
	}
}
