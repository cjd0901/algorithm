package _022_02

// Largest Number At Least Twice of Others
// leetcode: https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/

func LargestNumber(nums []int) int {
	maxIndex := 0
	for i, num := range nums {
		if num > nums[maxIndex] {
			maxIndex = i
		}
	}
	for _, num := range nums {
		max := nums[maxIndex]
		half := max / 2
		if num != max && half < num {
			return -1
		}
	}
	return maxIndex
}