package swapnodes

import (
	"reflect"
	"testing"
)

type Case struct {
	linkedList []int
	expected   []int
}

var cases = []Case{
	{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
	{[]int{1, 2, 3}, []int{2, 1, 3}},
	{[]int{}, []int{}},
	{[]int{5}, []int{5}},
	{[]int{1, 1, 5, 7, 9, 2, 23, 54, 1}, []int{1, 1, 7, 5, 2, 9, 54, 23, 1}},
}

func TestSwapPairs(t *testing.T) {
	for _, c := range cases {
		head := nodeFromList(c.linkedList)
		gotNode := swapPairs(head)
		got := listFromNode(gotNode)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"swapPairs(%v) == %v, want %v",
				c.linkedList,
				got,
				c.expected,
			)
		}
	}
}
