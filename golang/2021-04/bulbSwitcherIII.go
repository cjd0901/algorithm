package _021_04

// Bulb Switcher III
// leetCode: https://leetcode-cn.com/problems/bulb-switcher-iii/
// There is a room with n bulbs, numbered from 1 to n, arranged in a row from left to right.
// Initially, all the bulbs are turned off.
// At moment k (for k from 0 to n - 1), we turn on the light[k] bulb.
// A bulb change color to blue only if it is on and all the previous bulbs (to the left) are turned on too.
// Return the number of moments in which all turned on bulbs are blue.

func BulbSwitcherIII(light []int) int {
	curMax := 0
	ans := 0
	for i := 0; i < len(light); i++ {
		if curMax < light[i] {
			curMax = light[i]
		}
		if curMax == i+1 {
			ans ++
		}
	}
	return ans
}
