package groupthepeople

import (
	"reflect"
	"testing"
)

type Case struct {
	groupSizes []int
}

var cases = []Case{
	{[]int{3, 3, 3, 3, 3, 1, 3}},
	{[]int{2, 1, 3, 3, 3, 2}},
}

func TestGroupThePeople(t *testing.T) {
	for _, c := range cases {
		result := groupThePeople(c.groupSizes)

		// Based on the groups in the result, build up an input list
		compare := make([]int, len(c.groupSizes))
		var groupLen int
		for _, group := range result {
			groupLen = len(group)
			for _, index := range group {
				compare[index] = groupLen
			}
		}

		// Compare input list to the actual input
		bothEmpty := len(c.groupSizes) == 0 && len(compare) == 0
		if !reflect.DeepEqual(c.groupSizes, compare) && !bothEmpty {
			t.Errorf(
				"Invalid: groupThePeople(%v) == %v", c.groupSizes, result,
			)
		}
	}
}
