package getdecimalvalue

import (
	"testing"
)

// sliceToLinkedList converts a slice to a linked list and returns the head.
func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	node := head
	for i := 1; i < len(nums); i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		node.Next = newNode
		node = newNode
	}
	return head
}

type Case struct {
	nums     []int
	expected int
}

var cases = []Case{
	{[]int{1, 0, 1}, 5},
	{[]int{0}, 0},
	{[]int{1}, 1},
	{[]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}, 18880},
	{[]int{0, 0}, 0},
}

func TestGetDecimalValue(t *testing.T) {
	for _, c := range cases {
		head := sliceToLinkedList(c.nums)
		result := getDecimalValue(head)
		if result != c.expected {
			t.Errorf(
				"getDecimalValue(%v) == %v, want %v",
				c.nums, result, c.expected,
			)
		}
	}
}
