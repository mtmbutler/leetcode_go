package searchinsert

import "testing"

type Case struct {
	nums     []int
	target   int
	expected int
}

var cases = []Case{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0},
}

func TestSearchInsert(t *testing.T) {
	for _, c := range cases {
		got := searchInsert(c.nums, c.target)
		if got != c.expected {
			t.Errorf(
				"searchInsert(%v, %d) == %d, want %d",
				c.nums,
				c.target,
				got,
				c.expected,
			)
		}
	}
}
