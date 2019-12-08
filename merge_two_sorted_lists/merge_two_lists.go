// https://leetcode.com/problems/merge-two-sorted-lists/
package mergetwolists

// mergeTwoLists combines two sorted linked lists into a single sorted linked list by
// splicing together the nodes of the two given lists.
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Handle empty lists
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	// Put the main pointer on the lesser of the two head nodes. This main pointer
	// will represent the final linked list. The secondary pointer will stay on
	// whichever of the two original lists the main pointer is not pointing to, to make
	// sure we iterate through all the values.
	var head, main, other *ListNode
	if l1.Val <= l2.Val {
		main, other = l1, l2
	} else {
		main, other = l2, l1
	}
	head = main

	// Exhaust both lists
	for other != nil {
		if main.Next == nil {
			main.Next = other
			break
		} else if other.Val < main.Next.Val {
			main.Next, other = other, main.Next
		}
		main = main.Next
	}

	return head
}
