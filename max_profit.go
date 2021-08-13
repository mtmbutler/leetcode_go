// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
// https://leetcode.com/problems/maximum-subarray/
package main

import (
	"fmt"
)

type Case struct {
	prices   []int
	expected int
}

var cases = []Case{
	{[]int{7, 1, 5, 3, 6, 4}, 7},
	{[]int{1, 2, 3, 4, 5}, 4},
	{[]int{7, 6, 4, 3, 1}, 0},
	{[]int{8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6}, 15},
}

func main() {
	for _, c := range cases {
		got := maxProfit(c.prices)
		if got != c.expected {
			fmt.Printf(
				"maxProfit(%v) == %v, want %v\n",
				c.prices,
				got,
				c.expected,
			)
		}
	}
}

// maxProfit finds the maximum profit given a time series of profit.
// You are given an array prices where prices[i] is the price of a given stock on
// the ith day. Find the maximum profit you can achieve. You may complete as many
// transactions as you like (i.e., buy one and sell one share of the stock
// multiple times). Note: You may not engage in multiple transactions
// simultaneously (i.e., you must sell the stock before you buy again).
//
// The strategy here is to buy at every local minimum and sell at every local
// maximum.
func maxProfit(prices []int) int {
	profit := 0
	lastBuy := -1
	for i := 0; i < len(prices); i++ {
		if (i == 0 || prices[i] <= prices[i-1]) && (i == len(prices)-1 || prices[i] < prices[i+1]) && (lastBuy == -1) {
			// Local minimum; buy
			lastBuy = prices[i]
		} else if (i == 0 || prices[i] >= prices[i-1]) && (i == len(prices)-1 || prices[i] > prices[i+1]) && (lastBuy != -1) {
			// Local maximum; sell
			profit += prices[i] - lastBuy
			lastBuy = -1
		}
	}
	return profit
}
