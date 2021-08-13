// https://leetcode.com/problems/plus-one/
package main

import (
	"fmt"
	"reflect"
)

type Case struct {
	nums     []int
	expected []int
}

var cases = []Case{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	{[]int{4, 9}, []int{5, 0}},
	{[]int{9, 9, 9, 9}, []int{1, 0, 0, 0, 0}},
	{[]int{9, 3, 1, 9}, []int{9, 3, 2, 0}},
}

func main() {
	for _, c := range cases {
		got := plusOne(c.nums)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			fmt.Printf(
				"plusOne(%v) == %v, want %v\n",
				c.nums,
				got,
				c.expected,
			)
		}
	}
}

// plusOne adds one to an int represented as a slice.
// Given a non-empty array of decimal digits representing a non-negative integer,
// increment one to the integer. The digits are stored such that the most
// significant digit is at the head of the list, and each element in the array
// contains a single digit. You may assume the integer does not contain any
// leading zero, except the number 0 itself.
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// If we're here, the whole array was 9's
	return append([]int{1}, digits...)
}
