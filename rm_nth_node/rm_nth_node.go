// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package rmnthnode

// removeNthFromEnd removes the nth node from the end of a linked list and returns the
// head of the list. If n is greater than the length of the list, removeNthFromEnd
// will remove the first element.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Place a pointer at the head
	left := head

	// Place a second pointer n after the head. The idea is that, if we push the two
	// pointers along the linked list at the same rate, when the right pointer is
	// pointing at the last element, the left pointer will be one before the element we
	// want to remove. Then, we just point the left element to the one two after it,
	// and our job is done.
	right := head
	for i := 0; right != nil && i < n; i++ {
		right = right.Next
	}

	// If right is nil at this point, then the element we want to remove is the first
	// one, so let's just return the second node.
	if right == nil {
		return head.Next
	}

	// Push the pointers through the linked list until the right pointer is at the end
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	// Remove the unwanted element
	left.Next = left.Next.Next

	return head
}
