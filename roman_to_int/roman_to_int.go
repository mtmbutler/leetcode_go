// https://leetcode.com/problems/roman-to-integer/
package romantoint

var chars = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// romanToInt parses a string representing a valid Roman numeral into an integer.
func romanToInt(s string) int {
	num := 0
	last := 0
	for i := range s {
		x := chars[s[i]]

		// If the last number is smaller than the current, it needs to be subtracted
		// instead of added to the result. It's ok to do this at the first index, since
		// the last number is 0.
		if x > last {
			num -= 2 * last
		}
		num += x
		last = x
	}

	return num
}
