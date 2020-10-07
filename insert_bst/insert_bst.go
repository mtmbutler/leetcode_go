package insertbst

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// insertIntoBST inserts a given value into the given binary search tree.
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	for {
		if val > node.Val {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				return root
			}
			node = node.Right
		} else {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				return root
			}
			node = node.Left
		}

	}
}
