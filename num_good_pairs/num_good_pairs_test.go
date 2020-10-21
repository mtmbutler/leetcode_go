package numgoodpairs

import (
	"testing"
)

type Case struct {
	nums     []int
	expected int
}

var cases = []Case{
	{[]int{1, 2, 3, 1, 1, 3}, 4},
	{[]int{1, 1, 1, 1}, 6},
	{[]int{1, 2, 3}, 0},
}

func TestNumIdenticalPairs(t *testing.T) {
	for _, c := range cases {
		result := numIdenticalPairs(c.nums)
		if result != c.expected {
			t.Errorf(
				"numIdenticalPairs(%v) == %v, want %v",
				c.nums, result, c.expected,
			)
		}
	}
}
