package _021_05

// Bulb Switcher IV
// leetcode: https://leetcode-cn.com/problems/bulb-switcher-iv/

func BulbSwitcherIV(target string) int {
	count := 0
	cur := uint8(0)
	for i := 0; i < len(target); i++ {
		if target[i] - '0' != cur {
			count++
			cur = target[i] - '0'
		}
	}
	return count
}