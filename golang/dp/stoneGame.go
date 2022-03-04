package dp

// Stone Game
// leetcode: https://leetcode-cn.com/problems/stone-game/
type pair struct {
	first  int
	second int
}

func StoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]pair, n)
	for i := range dp {
		dp[i] = make([]pair, n)
		dp[i][i] = pair{
			first:  piles[i],
		}
	}

	for i := n-2; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			dp[i][j] = pair{}
			left := piles[i] + dp[i+1][j].second
			right := piles[j] + dp[i][j-1].second
			if left > right {
				dp[i][j].first = left
				dp[i][j].second = dp[i+1][j].first
			} else {
				dp[i][j].first = right
				dp[i][j].second = dp[i][j-1].first
			}
		}
	}
	return dp[0][n-1].first - dp[0][n-1].second > 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}