// https://leetcode.com/explore/challenge/card/august-leetcoding-challenge-2021/617/week-5-august-29th-august-31st/3957/
package main

import (
	"fmt"
)

type TestCase struct {
	m        int
	n        int
	ops      [][]int
	expected int
}

var testCases = []TestCase{
	{3, 3, [][]int{{2, 2}, {3, 3}}, 4},
	{3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}, 4},
	{3, 3, [][]int{}, 9},
	{3, 3, [][]int{{0, 0}}, 9},
	{3, 3, [][]int{{1, 2}, {2, 1}}, 1},
}

func main() {
	for _, c := range testCases {
		got := maxCount(c.m, c.n, c.ops)
		if got != c.expected {
			fmt.Printf(
				"maxCount(%v, %v, %v) == %v, want %v\n",
				c.m, c.n, c.ops, got, c.expected,
			)
		}
	}
}

// maxCount solves a LeetCode challenge called Range Addition II.
//
// Problem summary
//
// You are given an m x n matrix M initialized with all 0's and an array of
// operations ops, where ops[i] = [ai, bi] means M[x][y] should be incremented by
// one for all 0 <= x < ai and 0 <= y < bi.
//
// Count and return the number of maximum integers in the matrix after performing
// all the operations.
//
//
// Solution summary
//
// Since all the operations operate on a region starting with the top left
// element, the top left element will always be the max, and will increment with
// every non-zero operation. The trick, then, is to figure out how many cells are
// affected by every operation. This is done by finding the operation that affects
// the minumum number of rows, and the operation that affects the minimum number
// of columns (excluding operations that affect 0 cells). The size of the region
// affected by both of those operations is the number of elements that were
// incremented by every operation.
func maxCount(m int, n int, ops [][]int) int {
	minX := m
	minY := n
	for _, op := range ops {
		if op[0] == 0 || op[1] == 0 {
			// Operations with 0 in them don't affect any elements
			continue
		}
		if op[0] < minX {
			minX = op[0]
		}
		if op[1] < minY {
			minY = op[1]
		}
	}
	return minX * minY
}
