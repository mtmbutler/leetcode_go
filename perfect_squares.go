// https://leetcode.com/problems/perfect-squares/
package main

import (
	"fmt"
	"math"
)

type TestCase struct {
	n        int
	expected int
}

var cases = []TestCase{
	{0, 0},
	{1, 1},
	{2, 2},
	{4, 1},
	{12, 3},
	{13, 2},
	{68, 2},
	{144, 1},
}

func main() {
	for _, c := range cases {
		got := numSquares(c.n)
		if got != c.expected {
			fmt.Printf(
				"numSquares(%v) == %v, want %v\n",
				c.n,
				got,
				c.expected,
			)
		}
	}
}

// numSquares calculates the least number of perfect squares that sum to `n`.
//
// Given an integer n, return the least number of perfect square numbers that sum
// to n.
//
// A perfect square is an integer that is the square of an integer; in other
// words, it is the product of some integer with itself. For example, 1, 4, 9, and
// 16 are perfect squares while 3 and 11 are not.
func numSquares(n int) int {
	// Base cases
	if n <= 2 {
		return n
	}
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		return 1
	}

	// First collect the squares less than n
	squares := []int{}
	var sq int
	for k := 1; k < n; k++ {
		sq = k * k
		if sq > n {
			break
		}
		squares = append(squares, sq)
	}

	// BFS to find the quickest path to 0
	depth := 0
	queue := []int{n}
	var newQueue []int
	for len(queue) > 0 {
		depth++
		newQueue = []int{}
		for _, j := range queue {
			for _, square := range squares {
				if j == square {
					return depth
				}
				if j < square {
					break
				}
				newQueue = append(newQueue, j-square)
			}
		}
		queue = newQueue
	}
	return depth
}
