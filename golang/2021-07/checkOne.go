package _021_07

// Check If All 1's Are at Least Length K Places Away
// leetcode: https://leetcode-cn.com/problems/check-if-all-1s-are-at-least-length-k-places-away/
func KLengthApart(nums []int, k int) bool {
	preIndex := -1
	for i, n := range nums {
		if n == 1 {
			if preIndex == -1 || i - preIndex > k {
				preIndex = i
			} else {
				return false
			}
		}
	}
	return true
}