package kidswithcandies

import (
	"reflect"
	"testing"
)

type Case struct {
	candies      []int
	extraCandies int
	expected     []bool
}

var cases = []Case{
	{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
	{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	{[]int{}, 0, []bool{}},
}

func TestRunningSum(t *testing.T) {
	for _, c := range cases {
		got := kidsWithCandies(c.candies, c.extraCandies)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"kidsWithCandies(%v, %v) == %v, want %v",
				c.candies,
				c.extraCandies,
				got,
				c.expected,
			)
		}
	}
}
