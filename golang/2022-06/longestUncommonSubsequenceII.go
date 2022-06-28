package _022_06

// 最长特殊序列II
// leetcode: https://leetcode.cn/problems/longest-uncommon-subsequence-ii/
func LongestUncommonSubsequenceII(strs []string) int {
	ans := -1
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) <= ans {
			continue
		}
		valid := true
		for j := 0; j < len(strs); j++ {
			if i == j || !valid {
				continue
			}
			if isSubsequence(strs[i], strs[j]) {
				valid = false
			}
		}
		if valid {
			ans = max(ans, len(strs[i]))
		}
	}
	return ans
}

func isSubsequence(str1, str2 string) bool {
	l1, l2 := len(str1), len(str2)
	if l1 > l2 {
		return false
	}
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2] == l1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
