package golang

import "fmt"

// Best Time to Buy and Sell Stock with Transaction Fee
// leetCode: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i; and a non-negative integer fee representing a transaction fee.
// You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.
// You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)
// Return the maximum profit you can make.
// dynamic programming

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxStockProfit(prices []int, fee int) int {
	l := len(prices)
	s := make([][2]int, l)
	s[0][1] = -prices[0]
	for i := 1; i < l; i++ {
		s[i][0] = max(s[i-1][0], s[i-1][1]+prices[i]-fee)
		s[i][1] = max(s[i-1][1], s[i-1][0]-prices[i])
	}
	return s[l-1][0]
}

func MaxStockProfit2(prices []int, fee int) int {
	l := len(prices)
	sell := 0
	buy := -prices[0]
	for i := 1; i < l; i ++ {
		fmt.Println(sell, buy)
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, sell-prices[i])
		fmt.Println(sell, buy)
	}
	return sell
}