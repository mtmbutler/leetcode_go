package mergetwolists

import (
	"reflect"
	"testing"
)

type Case struct {
	l1       []int
	l2       []int
	expected []int
}

var cases = []Case{
	{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
	{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
	{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	{[]int{}, []int{}, []int{}},
}

func TestMergeTwoLists(t *testing.T) {
	for _, c := range cases {
		l1 := nodeFromList(c.l1)
		l2 := nodeFromList(c.l2)
		gotNode := mergeTwoLists(l1, l2)
		got := listFromNode(gotNode)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"mergeTwoLists(%v, %v) == %v, want %v",
				c.l1,
				c.l2,
				got,
				c.expected,
			)
		}
	}
}
