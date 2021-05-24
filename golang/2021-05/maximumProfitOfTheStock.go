package _021_05

// Maximum Profit of the Stock
// leetcode: https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

func MaximumProfitOfTheStock(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}
		maxProfit = max(maxProfit, prices[i] - minPrice)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}