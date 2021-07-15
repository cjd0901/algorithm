package _021_07

// Maximum Number of Balloons
// leetcode:https://leetcode-cn.com/problems/maximum-number-of-balloons/
func MaximumNumberOfBalloons(text string) int {
	target, ans := "balloon", len(text)
	standard := make(map[rune]int)
	letters := make([]int, 26)
	for _, c := range target {
		standard[c]++
	}
	for _, c := range text {
		letters[c-'a']++
	}
	for r, c := range standard {
		ans = min(ans, letters[r-'a']/c)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}