package getdecimalvalue

import (
	"testing"
)

type Case struct {
	nums     []int
	expected int
}

var cases = []Case{
	{[]int{1, 0, 1}, 5},
	{[]int{0}, 0},
	{[]int{1}, 1},
	{[]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}, 18880},
	{[]int{0, 0}, 0},
}

func TestGetDecimalValue(t *testing.T) {
	for _, c := range cases {
		head := nodeFromList(c.nums)
		result := getDecimalValue(head)
		if result != c.expected {
			t.Errorf(
				"getDecimalValue(%v) == %v, want %v",
				c.nums, result, c.expected,
			)
		}
	}
}
