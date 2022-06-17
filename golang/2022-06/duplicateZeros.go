package _022_06

// 复写零
// leetcode: https://leetcode.cn/problems/duplicate-zeros/
func DuplicateZeros(arr []int) {
	n := len(arr)
	l, r := -1, 0
	for r < n {
		l++
		if arr[l] == 0 {
			r += 2
		} else {
			r++
		}
	}
	j := n - 1
	if r == n+1 {
		arr[j] = 0
		j--
		l--
	}
	for j >= 0 {
		arr[j] = arr[l]
		j--
		if arr[l] == 0 {
			arr[j] = arr[l]
			j--
		}
		l--
	}
}
