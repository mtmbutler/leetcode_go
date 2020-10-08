// https://leetcode.com/problems/decoded-string-at-index/
package decodeatindex

import (
	"fmt"
	"strconv"
)

var digits = map[rune]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

// decodeAtIndex
// An encoded string S is given. To find and write the decoded string to a tape,
// the encoded string is read one character at a time and the following steps are
// taken:
//  - If the character read is a letter, that letter is written onto the tape.
//  - If the character read is a digit (say d), the entire current tape is
//    repeatedly written d-1 more times in total.
// Now for some encoded string S, and an index K, find and return the K-th letter
// (1 indexed) in the decoded string.
func decodeAtIndex(S string, K int) string {
	tape := []rune{}
	for i, char := range S {
		_, err := strconv.Atoi(string(S[i+1:]))
		if err == nil && len(tape) >= K {
			break
		}
		digit, ok := digits[char]
		if ok {
			// If we have a digit, repeat the existing slice
			for j := 1; j < digit; j++ {
				for _, existingChar := range tape {
					tape = append(tape, existingChar)
				}
			}
		} else {
			// If we don't have a digit, just add the character to the tape
			tape = append(tape, char)
		}

	}
	fmt.Println(string(tape))
	return string(tape[K-1])
}
