package maxdepth

// maxDepth determines the maximum parentheses depth of a valid parentheses
// expression.
func maxDepth(s string) int {
	depth := 0
	maxDepth := 0
	for _, ch := range s {
		if ch == '(' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if ch == ')' {
			depth--
		}
	}
	return maxDepth
}
