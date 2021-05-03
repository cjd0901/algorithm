package _021_05

// Predict The Winner
// leetcode: https://leetcode-cn.com/problems/predict-the-winner/
// You are given an integer array nums.
// Two players are playing a game with this array: player 1 and player 2.
// Player 1 and player 2 take turns, with player 1 starting first.
// Both players start the game with a score of 0.
// At each turn, the player takes one of the numbers from either end of the array (i.e., nums[0] or nums[nums.length - 1]) which reduces the size of the array by 1.
// The player adds the chosen number to their score.
// The game ends when there are no more elements in the array.
// Return true if Player 1 can win the game.
// If the scores of both players are equal, then player 1 is still the winner, and you should also return true.
// You may assume that both players are playing optimally.

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}