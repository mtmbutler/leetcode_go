// https://leetcode.com/problems/maximum-subarray/
package main

import (
	"fmt"
)

type Case struct {
	nums     []int
	expected int
}

var cases = []Case{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{1}, 1},
	{[]int{-1}, -1},
	{[]int{5, 4, -1, 7, 8}, 23},
}

func main() {
	for _, c := range cases {
		got := maxSubArray(c.nums)
		if got != c.expected {
			fmt.Printf(
				"maxSubArray(%v) == %v, want %v\n",
				c.nums,
				got,
				c.expected,
			)
		}
	}
}

// maxSubArray finds the largest sum of a contiguous subarray in the
// provided list of integers.
// This is a brute force implementation that considers every contiguous
// subarray by means of a window.
func maxSubArray(nums []int) int {
	var largestSum int = nums[0]
	arrLen := len(nums)
	for leftIndex := 0; leftIndex < arrLen; leftIndex++ {
		for rightIndex := leftIndex + 1; rightIndex <= arrLen; rightIndex++ {
			subArraySum := sum(nums[leftIndex:rightIndex])
			if subArraySum > largestSum {
				largestSum = subArraySum
			}
		}
	}
	return largestSum
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}
