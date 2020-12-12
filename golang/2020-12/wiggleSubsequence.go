package golang

// Wiggle Subsequence
// leetCode: https://leetcode-cn.com/problems/wiggle-subsequence/

func WiggleSubsequence(nums []int) (I int) {
	var previous string
	if len(nums) < 2 {
		return len(nums)
	}
	I = 1
	for i := 0; i < len(nums); i++ {
		if i == len(nums) - 1 {
			break
		}
		if nums[i] - nums[i+1] < 0 && previous != "N" {
			previous = "N"
			I++
		}else if nums[i] - nums[i+1] > 0 && previous != "P" {
			previous = "P"
			I++
		}
	}
	return
}