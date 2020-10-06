package runningsumarray

import (
	"reflect"
	"testing"
)

type Case struct {
	nums     []int
	expected []int
}

var cases = []Case{
	{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
	{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	{[]int{}, []int{}},
}

func TestRunningSum(t *testing.T) {
	for _, c := range cases {
		got := runningSum(c.nums)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"runningSum(%v) == %v, want %v",
				c.nums,
				got,
				c.expected,
			)
		}
	}
}
