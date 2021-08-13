// https://leetcode.com/problems/plus-one/
package main

import (
	"fmt"
)

type Case struct {
	x        int
	expected int
}

var cases = []Case{
	{4, 2},
	{8, 2},
	{160, 12},
	{0, 0},
	{1, 1},
	{3, 1},
	{2, 1},
}

func main() {
	for _, c := range cases {
		got := mySqrt(c.x)
		if got != c.expected {
			fmt.Printf(
				"mySqrt(%v) == %v, want %v\n",
				c.x,
				got,
				c.expected,
			)
		}
	}
}

// mySqrt returns the integer square root of an integer, rounded down.
// Given a non-negative integer x, compute and return the square root of x. Since
// the return type is an integer, the decimal digits are truncated, and only the
// integer part of the result is returned. Note: You are not allowed to use any
// built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.
//
// This implementation uses the Newton-Raphson method.
func mySqrt(x int) int {
	guess := x
	for guess*guess > x {
		guess = (guess + x/guess) / 2
	}
	return guess
}
