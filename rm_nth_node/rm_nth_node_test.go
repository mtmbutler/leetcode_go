package rmnthnode

import (
	"reflect"
	"testing"
)

type Case struct {
	list     []int
	n        int
	expected []int
}

var cases = []Case{
	{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
	{[]int{1, 2}, 1, []int{1}},
	{[]int{2, 5, 4, 7, 8}, 3, []int{2, 5, 7, 8}},
	{[]int{1}, 1, nil},
	{[]int{1, 2}, 2, []int{2}},
	{[]int{1, 2}, 45, []int{2}},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for _, c := range cases {
		head := nodeFromList(c.list)
		gotNode := removeNthFromEnd(head, c.n)
		got := listFromNode(gotNode)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"removeNthFromEnd(%v, %d) == %v, want %v",
				c.list,
				c.n,
				got,
				c.expected,
			)
		}
	}
}
