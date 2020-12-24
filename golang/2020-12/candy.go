package golang

// Candy
// leetCode: https://leetcode-cn.com/problems/candy/
// There are N children standing in a line. Each child is assigned a rating value.
// You are giving candies to these children subjected to the following requirements:
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// What is the minimum candies you must give?

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Candy(ratings []int) (I int) {
	n := len(ratings)
	left := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}else {
			left[i] = 1
		}
	}
	right := 0
	for i := n-1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		}else {
			right = 1
		}
		I += max(left[i], right)
	}
	return
}
