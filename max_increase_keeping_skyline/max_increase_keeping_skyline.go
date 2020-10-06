// https://leetcode.com/problems/max-increase-to-keep-city-skyline/
package maxincreasekeepingskyline

// maxIncreaseKeepingSkyline determines, for a grid of buildings, the maximum
// total height increase possible without changing the skyline from the edges.
func maxIncreaseKeepingSkyline(grid [][]int) int {
	// Determine the tallest building in each row and column
	rowMaxes := make([]int, len(grid))
	var colMaxes []int
	for i, row := range grid {
		if colMaxes == nil {
			colMaxes = make([]int, len(row))
		}
		for j, height := range row {
			if height > rowMaxes[i] {
				rowMaxes[i] = height
			}
			if height > colMaxes[j] {
				colMaxes[j] = height
			}
		}
	}

	// Determine the max increase for each building, and keep a running total
	total := 0
	for i, row := range grid {
		for j, height := range row {
			var ceiling int
			if rowMaxes[i] < colMaxes[j] {
				ceiling = rowMaxes[i]
			} else {
				ceiling = colMaxes[j]
			}
			if ceiling > height {
				total += ceiling - height
			}
		}
	}
	return total
}
