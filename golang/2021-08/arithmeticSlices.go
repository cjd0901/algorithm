package _021_08

// Arithmetic Slices
// leetcode: https://leetcode-cn.com/problems/arithmetic-slices/
func ArithmeticSlices(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n-2; {
		j, d := i, nums[i+1] - nums[i]
		for j+1 < n && nums[j+1] - nums[j] == d {
			j++
		}
		l := j-i+1
		s, e := 1, l-3+1
		ans += (s+e) * e /2
		i = j
	}
	return ans
}

func ArithmeticSlices2(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := 0
	d, t := nums[1]-nums[0], 0
	for i := 2; i < n; i++ {
		if nums[i] - nums[i-1] == d {
			t++
		}else {
			d, t = nums[i] - nums[i-1], 0
		}
		ans += t
	}
	return ans
}

func ArithmeticSlices3(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i + 1 < n; i++ {
		if nums[i]*2 == nums[i-1] + nums[i+1] {
			dp[i] = dp[i-1] + 1
		}
	}
	sum := 0
	for _, c := range dp {
		sum += c
	}
	return sum
}