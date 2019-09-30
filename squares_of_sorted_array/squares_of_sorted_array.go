// https://leetcode.com/problems/squares-of-a-sorted-array/
package squaresofsortedarray

// sortedSquares returns a sorted array of the element-wise squares of the input array.
func sortedSquares(A []int) []int {
	// Keep two arrays -- one for the squares of the negative elements, and another for
	// the squares of the positive. The first will be descending, and the second
	// ascending, and we'll merge them together for the final array.
	var neg []int
	var pos []int
	for _, i := range A {
		if i < 0 {
			neg = append(neg, i*i)
		} else {
			pos = append(pos, i*i)
		}
	}

	// Merge the arrays by keeping a pointer on each
	var finalArr []int
	var next int
	var noMoreNeg bool
	var noMorePos bool
	i := 0
	j := len(neg) - 1
	for {
		noMoreNeg = j < 0
		noMorePos = i >= len(pos)
		if noMorePos && noMoreNeg {
			break
		} else if noMorePos {
			next = neg[j]
			j--
		} else if noMoreNeg {
			next = pos[i]
			i++
		} else if pos[i] <= neg[j] {
			next = pos[i]
			i++
		} else {
			next = neg[j]
			j--
		}
		finalArr = append(finalArr, next)
	}

	return finalArr
}
