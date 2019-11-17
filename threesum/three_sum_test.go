package threesum

import (
	"reflect"
	"testing"
)

type Case struct {
	nums     []int
	expected [][]int
}

var cases = []Case{
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
}

func TestThreeSum(t *testing.T) {
	for _, c := range cases {
		got := threeSum(c.nums)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"threeSum(%v) == %v, want %v",
				c.nums,
				got,
				c.expected,
			)
		}
	}
}
