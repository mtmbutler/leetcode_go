package twosum

import (
	"reflect"
	"testing"
)

type Case struct {
	nums     []int
	target   int
	expected []int
}

var cases = []Case{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{2, 0, 2, 1}, 4, []int{0, 2}},
}

func TestTwoSum(t *testing.T) {
	for _, c := range cases {
		got := twoSum(c.nums, c.target)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"twoSum(%v, %d) == %v, want %v",
				c.nums,
				c.target,
				got,
				c.expected,
			)
		}
	}
}
