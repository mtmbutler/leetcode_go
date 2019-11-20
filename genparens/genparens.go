package genparens

func generateParenthesis(n int) []string {
	var left, surround, right string
	if n <= 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	} else {
		// Generate the combinations recursively
		// Add to a set to prevent duplicates
		set := map[string]struct{}{}
		base := generateParenthesis(n - 1)
		for _, validParens := range base {
			left = "()" + validParens
			surround = "(" + validParens + ")"
			right = validParens + "()"
			set[left] = struct{}{}
			set[surround] = struct{}{}
			set[right] = struct{}{}
		}

		// Convert the set to a slice
		slc := []string{}
		for s, _ := range set {
			slc = append(slc, s)
		}
		return slc
	}
}
