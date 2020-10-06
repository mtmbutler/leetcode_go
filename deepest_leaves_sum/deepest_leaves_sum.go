package deepestleaves

// TreeNode is a node of a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// deepestLeavesSum calculates the sum of all the deepest nodes of a tree. This
// is implemented with a breadth-first search. Whenever we reach a new level of
// depth, we reset the sum to 0.
func deepestLeavesSum(root *TreeNode) int {
	var sum int
	currentLevel := []*TreeNode{root}
	for len(currentLevel) != 0 {
		sum = 0
		nextLevel := []*TreeNode{}
		for _, node := range currentLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
	}
	return sum
}
