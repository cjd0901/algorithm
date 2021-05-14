package _021_05

// Minimum Elements to Add to Form a Given Sum
// leetcode: https://leetcode-cn.com/problems/minimum-elements-to-add-to-form-a-given-sum/
// You are given an integer array nums and two integers limit and goal.
// The array nums has an interesting property that abs(nums[i]) <= limit.
// Return the minimum number of elements you need to add to make the sum of the array equal to goal.
// The array must maintain its property that abs(nums[i]) <= limit.

func MinimumElementsToAddToFormAGivenSum(nums []int, limit, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	diff := abs(sum - goal)
	if diff % limit == 0 {
		return diff / limit
	}else {
		return diff / limit + 1
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}