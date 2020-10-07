package dividetwointegers

import "testing"

type Case struct {
	dividend int
	divisor  int
	expected int
}

var cases = []Case{
	{10, 3, 3},
	{7, -3, -2},
	{0, 1, 0},
	{1, 1, 1},
	{50, 5, 10},
	{100, 21, 4},
	{0, -1, 0},
	{-1, -1, 1},
}

func TestDivide(t *testing.T) {
	for _, c := range cases {
		result := divide(c.dividend, c.divisor)
		if result != c.expected {
			t.Errorf(
				"divide(%v, %v) == %v, want %v",
				c.dividend,
				c.divisor,
				result,
				c.expected,
			)
		}
	}
}
