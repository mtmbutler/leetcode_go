// https://leetcode.com/problems/two-sum/
package twosum

// twoSum finds the indices of two numbers in the given list such that the numbers add
// up to the given target. It is assumed that the given list has exactly one solution.
// Behavior in other cases is undefined. twoSum won't use the same index twice.
func twoSum(nums []int, target int) []int {
	// First, put the list into a hash of value -> [ix1, ix2, etc.]
	hash := make(map[int][]int)
	for i, x := range nums {
		_, ok := hash[x]
		if ok {
			hash[x] = append(hash[x], i)
		} else {
			hash[x] = []int{i}
		}
	}

	// Now, loop through the list again. For each element, calculate the difference
	// between the value and the target, and look up that difference in the hash. If it
	// exists at a different index than the one we're at, return the current index and
	// that index.
	for i, x := range nums {
		diff := target - x
		ixs, ok := hash[diff]
		if ok {
			if diff == x && len(ixs) > 1 {
				return []int{i, ixs[1]}
			} else if diff != x {
				return []int{i, ixs[0]}
			}
		}
	}

	return []int{}
}
