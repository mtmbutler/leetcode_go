// https://leetcode.com/problems/3sum/
package threesum

import "sort"

// threeSum finds all unique triplets of array elements such that the
// sum of each triplet is 0.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := [][]int{}
	lastIx := len(nums) - 1
	var target, twoSum, i, j, k, x, y, z int

	for i = 0; i < lastIx-1; i++ {
		x = nums[i]
		target = -x
		j = i + 1
		k = lastIx
		for j < k {
			y = nums[j]
			z = nums[k]
			twoSum = y + z
			if twoSum < target {
				j++
			} else if twoSum > target {
				k--
			} else {
				triplets = append(triplets, []int{x, y, z})
				for j < k && nums[j] == y {
					j++
				}
				for j < k && nums[k] == z {
					k--
				}
			}
		}
		for i < lastIx && nums[i] == nums[i+1] {
			i++
		}
	}

	return triplets
}
