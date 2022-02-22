package _022_02

// Pancake Sorting
// leetcode: https://leetcode-cn.com/problems/pancake-sorting/
func PancakeSorting(arr []int) []int {
	ans := make([]int, 0)
	for n := len(arr); n > 1; n-- {
		idx := 0
		for i, v := range arr[:n] {
			if v > arr[idx] {
				idx = i
			}
		}
		if idx == n-1 {
			continue
		}
		for i := 0; i < (idx+1)/2; i++ {
			arr[i], arr[idx-i] = arr[idx-i], arr[i]
		}
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
		}
		ans = append(ans, idx+1, n)
	}
	return ans
}