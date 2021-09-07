// https://leetcode.com/problems/smallest-string-with-swaps/
package main

import (
	"fmt"
	"strings"
)

type Case struct {
	s        string
	pairs    [][]int
	expected string
}

var cases = []Case{
	{s: "dcab", pairs: [][]int{{0, 3}, {1, 2}}, expected: "bacd"},
	{s: "dcab", pairs: [][]int{{0, 3}, {1, 2}, {0, 2}}, expected: "abcd"},
	{s: "cba", pairs: [][]int{{0, 1}, {1, 2}}, expected: "abc"},
	{s: "otilzqqoj", pairs: [][]int{{2, 3}, {7, 3}, {3, 8}, {1, 7}, {1, 0}, {0, 4}, {0, 6}, {3, 4}, {2, 5}}, expected: "ijlooqqtz"},
	{s: "vbjjxgdfnru", pairs: [][]int{{8, 6}, {3, 4}, {5, 2}, {5, 5}, {3, 5}, {7, 10}, {6, 0}, {10, 0}, {7, 1}, {4, 8}, {6, 2}}, expected: "bdfgjjnruvx"},
}

func main() {
	for _, c := range cases {
		got := smallestStringWithSwaps(c.s, c.pairs)
		if got != c.expected {
			fmt.Printf(
				"smallestStringWithSwaps(%v, %v) == %v, want %v\n",
				c.s,
				c.pairs,
				got,
				c.expected,
			)
		}
	}
}

// swap returns a modified string with the characters at the indices indicated by
// pair swapped.
func swap(s string, pair []int) string {
	// Collect bytes for new string
	i, j := pair[0], pair[1]
	bytes := []byte{}
	for k := range s {
		if k == i {
			bytes = append(bytes, s[j])
		} else if k == j {
			bytes = append(bytes, s[i])
		} else {
			bytes = append(bytes, s[k])
		}
	}

	// Build string
	builder := strings.Builder{}
	builder.Write(bytes)
	return builder.String()
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	candidates := []string{s}
	queue := []string{s}
	checked := map[string]struct{}{s: {}}
	var newQueue []string
	var smallerFound bool
	var newStr string
	for len(queue) > 0 {
		newQueue = []string{}
		for _, st := range queue {
			smallerFound = false
			for _, pair := range pairs {
				newStr = swap(st, pair)
				if _, ok := checked[newStr]; !ok {
					checked[newStr] = struct{}{}
					if newStr < st {
						smallerFound = true
					}
					newQueue = append(newQueue, newStr)
				}
			}
			if !smallerFound {
				candidates = append(candidates, st)
			}
		}
		queue = newQueue
	}

	for _, st := range candidates {
		if st < s {
			s = st
		}
	}
	return s
}
