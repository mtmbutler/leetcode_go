package strstr

import "testing"

type Case struct {
	haystack string
	needle   string
	expected int
}

func TestStrstr(t *testing.T) {
	// cases := [
	// Case{"hello", "ll", 2}]
	for _, c := range cases {
		got := strStr(c.haystack, c.needle)
		if got != c.want {
			t.Errorf(
				"strStr(%q, %q) == %q, want %q",
				c.haystack,
				c.needle,
				c.expected,
				c.want,
			)
		}
	}
	got := strStr()
	want := 1
	if got != want {
		t.Errorf("hello() == %q, want %q", got, want)
	}
}
