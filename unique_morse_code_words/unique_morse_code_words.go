// https://leetcode.com/problems/unique-morse-code-words/
package uniquemorsecodewords

import "strings"

var morse = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
}

// uniqueMorseRepresentations determines the morse representations of the words passed,
// then counts the number of unique representations in the resulting list.
func uniqueMorseRepresentations(words []string) int {
	// Instantiate an empty set to add representations to
	set := map[string]struct{}{}

	for _, word := range words {
		var strBuilder strings.Builder
		// 1. Convert the word to morse code
		for _, c := range word {
			strBuilder.WriteString(morse[c])
		}

		// 2. Add the transformed word to the set
		morseRep := strBuilder.String()
		set[morseRep] = struct{}{}
	}
	return len(set)
}
