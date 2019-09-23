package strstr

import "testing"

type Case struct {
	haystack string
	needle   string
	expected int
}

func TestStrstr(t *testing.T) {
	cases := []Case{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "a", -1},
		{"hello", "", 0},
		{"", "", 0},
	}
	for _, c := range cases {
		got := strStr(c.haystack, c.needle)
		if got != c.expected {
			t.Errorf(
				"strStr(%q, %q) == %d, want %d",
				c.haystack,
				c.needle,
				got,
				c.expected,
			)
		}
	}
}
