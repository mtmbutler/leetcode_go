// https://leetcode.com/problems/running-sum-of-1d-array/
package runningsumarray

// runningSum calculates a running sum of the input list.
func runningSum(nums []int) []int {
	s := 0
	rs := []int{}
	for _, x := range nums {
		s += x
		rs = append(rs, s)
	}
	return rs
}
