package containermostwater

import "testing"

type Case struct {
	height   []int
	expected int
}

var cases = []Case{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{0, 0}, 0},
	{[]int{1, 5, 6, 9}, 10},
	{[]int{2, 3}, 2},
}

func TestMaxArea(t *testing.T) {
	for _, c := range cases {
		got := maxArea(c.height)
		if got != c.expected {
			t.Errorf(
				"maxArea(%v) == %d, want %d",
				c.height,
				got,
				c.expected,
			)
		}
	}
}
