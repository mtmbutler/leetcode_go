package getdecimalvalue

import "math"

// getDecimalValue ...
// Given head which is a reference node to a singly-linked list. The value of each
// node in the linked list is either 0 or 1. The linked list holds the binary
// representation of a number.
// Return the decimal value of the number in the linked list.
func getDecimalValue(head *ListNode) int {
	// Load the linked list into a slice
	nums := []int{head.Val}
	for head.Next != nil {
		head = head.Next
		nums = append(nums, head.Val)
	}

	// Use positional method to calculate decimal
	dec := 0
	for i, n := range nums {
		power := len(nums) - i - 1
		dec += n * int(math.Exp2(float64(power)))
	}

	return dec
}
