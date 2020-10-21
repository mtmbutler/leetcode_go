package shuffle

import (
	"reflect"
	"testing"
)

type Case struct {
	nums     []int
	n        int
	expected []int
}

var cases = []Case{
	{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
	{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
	{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
}

func TestShuffle(t *testing.T) {
	for _, c := range cases {
		result := shuffle(c.nums, c.n)
		bothEmpty := len(c.expected) == 0 && len(result) == 0
		if !bothEmpty && !reflect.DeepEqual(c.expected, result) {
			t.Errorf(
				"shuffle(%v, %v) == %v, want %v",
				c.nums, c.n, result, c.expected,
			)
		}
	}
}
