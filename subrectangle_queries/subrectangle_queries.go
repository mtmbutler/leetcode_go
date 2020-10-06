// https://leetcode.com/problems/subrectangle-queries/
package subrectanglequeries

// SubrectangleQueries stores a matrix and allows some operations.
type SubrectangleQueries struct {
	arr [][]int
}

// Constructor creates a new SubrectangleQueries struct
func Constructor(rectangle [][]int) SubrectangleQueries {
	sq := SubrectangleQueries{}
	sq.arr = rectangle
	return sq
}

// UpdateSubrectangle applies newValue to every cell in the given matrix between
// top-left corner (row1, col1) and bottom-right corner (row2, col2).
func (this *SubrectangleQueries) UpdateSubrectangle(
	row1 int, col1 int, row2 int, col2 int, newValue int,
) {
	for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			this.arr[r][c] = newValue
		}
	}
}

// GetValue retrieves the cell at (row, col)
func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.arr[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
