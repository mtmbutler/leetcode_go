package searchinsert

// searchInsert returns the index of `target` in `nums`, a sorted array with no
// duplicates. If `target` is not found, returns the index where it would be inserted
// in order.
func searchInsert(nums []int, target int) int {
	for i, n := range nums {
		if n >= target {
			return i
		}
	}
	return len(nums)
}
