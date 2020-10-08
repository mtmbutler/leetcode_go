package decodeatindex

import "testing"

type Case struct {
	S        string
	K        int
	expected string
}

var cases = []Case{
	{"leet2code3", 10, "o"},
	{"ha22", 5, "h"},
	{"a2345678999999999999999", 1, "a"},
}

func TestDecodeAtIndex(t *testing.T) {
	for _, c := range cases {
		result := decodeAtIndex(c.S, c.K)
		if result != c.expected {
			t.Errorf(
				"decodeAtIndex(%v, %v) == %v, want %v",
				c.S, c.K, result, c.expected,
			)
		}
	}
}
