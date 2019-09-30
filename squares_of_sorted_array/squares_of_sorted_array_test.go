package squaresofsortedarray

import (
	"reflect"
	"testing"
)

type Case struct {
	arr      []int
	expected []int
}

var cases = []Case{
	{[]int{}, []int{}},
	{[]int{1, 2}, []int{1, 4}},
	{[]int{5}, []int{25}},
	{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
}

func TestSortedSquares(t *testing.T) {
	for _, c := range cases {
		got := sortedSquares(c.arr)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"sortedSquares(%v) == %v, want %v",
				c.arr,
				got,
				c.expected,
			)
		}
	}
}
