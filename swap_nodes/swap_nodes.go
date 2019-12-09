// https://leetcode.com/problems/swap-nodes-in-pairs/
package swapnodes

// swapPairs swaps every two adjacent nodes in a linked list.
// Example:
//   Input: 1 -> 2 -> 3 -> 4
//   Output: 2 -> 1 -> 4 -> 3
func swapPairs(head *ListNode) *ListNode {
	// Put a phantom node before the start so we don't have to do any special handling
	// for the first pair of nodes
	phantomNode := ListNode{0, head}
	left, right := &phantomNode, head

	// Use two pointers and a six-step cycle to iterate through the list and swap each
	// pair of two nodes.
	for left.Next != nil && right.Next != nil {
		left.Next = right.Next
		left = right
		right = left.Next
		left.Next = right.Next
		right.Next = left
		right = left.Next
	}

	return phantomNode.Next
}
