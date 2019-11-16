// https://leetcode.com/problems/zigzag-conversion/
package zigzagconversion

import "strings"

// convert writes the given string in a zigzag pattern over the specified number of
// rows, then returns it as read line-by-line.
//
// Example:
//     s = "PAYPALISHIRING", numRows = 3
//
//     First, we write the string in zigzag form:
//	     P   A   H   N
//		 A P L S I I G
//		 Y	 I   R
//
//     Then, we return it read line-by-line:
//       PAHNAPLSIIGYIR
func convert(s string, numRows int) string {
	rows := make([]strings.Builder, numRows)
	var strBuilder strings.Builder
	var ch string // Holds the character to add to the builder

	ix := 0
	asc := true // Whether to increase the row index after this character
	for _, c := range s {
		ch = string(c)
		rows[ix].WriteString(ch)
		if ix == 0 {
			asc = true // If we're at the first row, switch to ascending
		} else if ix == numRows-1 {
			asc = false // If we're at the last row, switch to descending
		}

		// Increment/decrement the index, unless there's only one row
		if numRows == 1 {

		} else if asc {
			ix++
		} else {
			ix--
		}
	}

	// Read each row into a new string
	for _, b := range rows {
		ch = b.String()
		strBuilder.WriteString(ch)
	}
	combined := strBuilder.String()

	return combined
}
