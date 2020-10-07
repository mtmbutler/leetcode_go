package maxbinarytree

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// constructMaximumBinaryTree
// Given an integer array with no duplicates. A maximum tree building on this
// array is defined as follows:
//  - The root is the maximum number in the array.
//  - The left subtree is the maximum tree constructed from left part subarray
//    divided by the maximum number.
//  - The right subtree is the maximum tree constructed from right part subarray
//    divided by the maximum number.
// Construct the maximum tree by the given array and output the root node of this
// tree.
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var max int
	var maxIndex int
	for i, x := range nums {
		if i == 0 || x > max {
			max = x
			maxIndex = i
		}
	}
	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}
