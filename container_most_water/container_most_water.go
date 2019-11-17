// https://leetcode.com/problems/container-with-most-water/
package containermostwater

// min takes the lower of two integers
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// max takes the greater of two integers
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// maxArea takes an array of non-negative integers a_1, a_2, ..., a_n, where each
// element represents a barrier of height a_i at coordinate i, and calculates the
// maximum area of water that can be stored between two barriers in the array.
func maxArea(height []int) int {
	var area int
	maxArea := -1
	for i, a := range height {
		for j, b := range height {
			if j <= i {
				continue
			} else {
				area = min(a, b) * (j - i)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}
