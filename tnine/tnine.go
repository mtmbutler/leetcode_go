// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
package tnine

var numberPad = map[rune][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

// letterCombinations returns all the possible letter combinations that the provided
// number could represent, if it were typed into a phone keypad.
func letterCombinations(digits string) []string {
	combinations := []string{}
	var combination string
	var prevCombinations []string
	var letters []string

	for _, ch := range digits {
		prevCombinations = combinations
		letters = numberPad[ch]
		if len(prevCombinations) == 0 {
			combinations = letters
		} else {
			combinations = []string{}
			for _, c := range prevCombinations {
				for _, letter := range letters {
					combination = c + letter
					combinations = append(combinations, combination)
				}
			}

		}
	}

	return combinations
}
