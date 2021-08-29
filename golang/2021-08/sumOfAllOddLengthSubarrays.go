package _021_08

// Sum of All Odd Length Subarrays
// leetcode: https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/
func SumOfAllOddLengthSubarrays(arr []int) int {
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if (j - i + 1) % 2 == 1 {
				if i == 0 {
					ans += arr[j]
				} else {
					ans += arr[j] - arr[i-1]
				}
			}
		}
	}
	return ans
}