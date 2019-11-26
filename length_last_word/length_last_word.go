package lenlastword

// lengthOfLastWord counts the length of the final word that appears in the string. If
// the string is empty or has all spaces, return 0. If there are no spaces, the whole
// string is considered one word, and the length of the string will be returned.
func lengthOfLastWord(s string) int {
	var isSpace bool  // Whether the character is a space
	counter := 0      // Track the length of the last word
	readWord := false // Whether the lookback has encountered a non-space yet
	for i := len(s) - 1; i >= 0; i-- {
		isSpace = s[i] == ' '
		if isSpace && readWord {
			break
		} else if !isSpace {
			readWord = true
			counter++
		}
	}
	return counter
}
