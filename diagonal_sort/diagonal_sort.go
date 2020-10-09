package diagsort

import (
	"sort"
)

// Index represents a point in a 2-D matrix. I = row, J = col.
type Index struct {
	I, J int
}

// diagonalSort sorts each top-left to bottom-right diagonal in a matrix.
func diagonalSort(mat [][]int) [][]int {
	// Determine dimensions and populate the result matrix
	m := len(mat)
	n := 0
	if m > 0 {
		n = len(mat[0])
	}
	result := make([][]int, m)
	for row := range result {
		result[row] = make([]int, n)
	}

	// Capture each diagonal in a map of (top left point): [values]
	// For each point, the top left point can be calculated by subtracting the
	// lesser of i and j from both of them.
	diags := map[Index][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diffToZero := i
			if j < i {
				diffToZero = j
			}
			startingPoint := Index{i - diffToZero, j - diffToZero}
			diags[startingPoint] = append(diags[startingPoint], mat[i][j])
		}
	}

	// Sort each diagonal and update the result matrix
	for startingPoint, diag := range diags {
		sort.Ints(diag)
		i := startingPoint.I
		j := startingPoint.J
		for _, val := range diag {
			result[i][j] = val
			i++
			j++
		}
	}

	return result
}
