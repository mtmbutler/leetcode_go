// https://leetcode.com/problems/implement-strstr/
package strstr

// strStr finds the first index where substring `needle` occurs in string `haystack`.
// If needle is the empty string, returns 0. If needle doesn't occur in haystack,
// returns -1.
func strStr(haystack, needle string) int {
	// Store lengths of the strings
	haystackLen := len(haystack)
	needleLen := len(needle)

	// Handle edge cases where the needle is empty, equal to or longer than the
	// haystack
	if needleLen == 0 {
		return 0
	} else if needle == haystack {
		return 0
	} else if needleLen > haystackLen {
		return -1
	}

	// Iterate through the haystack, checking against the needle
	lastIndex := haystackLen - needleLen
	for i := 0; i <= lastIndex; i++ {
		window := haystack[i : i+needleLen]
		if window == needle {
			return i
		}
	}
	return -1
}
