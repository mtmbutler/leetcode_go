package swapnodes

// Leetcode's definition of a linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// nodeFromList returns the head of a linked list representing the passed list
func nodeFromList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{arr[0], nil}
	last := head
	for _, x := range arr[1:] {
		newNode := &ListNode{x, nil}
		last.Next = newNode
		last = newNode
	}
	return head
}

// listfromNode returns a slice representing the passed linked list
func listFromNode(node *ListNode) []int {
	li := []int{}
	if node != nil {
		for {
			li = append(li, node.Val)
			if node.Next == nil {
				break
			} else {
				node = node.Next
			}
		}
	}
	return li
}
