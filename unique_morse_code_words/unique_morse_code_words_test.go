package uniquemorsecodewords

import "testing"

type Case struct {
	words    []string
	expected int
}

var cases = []Case{
	{[]string{"foo", "bar"}, 2},
	{[]string{"gin", "zen", "gig", "msg"}, 2},
}

func TestUniqueMorseRepresentations(t *testing.T) {
	for _, c := range cases {
		got := uniqueMorseRepresentations(c.words)
		if got != c.expected {
			t.Errorf(
				"uniqueMorseRepresentations(%s) == %d, want %d",
				c.words,
				got,
				c.expected,
			)
		}
	}
}
