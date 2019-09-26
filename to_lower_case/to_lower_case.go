// https://leetcode.com/problems/to-lower-case/
package tolowercase

import "strings"

// toLowerCase converts an input string to lowercase.
func toLowerCase(str string) string {
	var strBuilder strings.Builder
	var r rune
	var ch string // Holds the character to add to the builder

	for _, c := range str {
		if 'A' <= c && c <= 'Z' {
			r = c - 'A' + 'a'
			ch = string(r)
		} else {
			ch = string(c)
		}
		strBuilder.WriteString(ch)
	}
	lowerStr := strBuilder.String()
	return lowerStr
}
