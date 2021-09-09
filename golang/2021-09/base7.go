package _021_09

import "fmt"

// Base 7
// leetcode: https://leetcode-cn.com/problems/base-7/

func Base7(num int) string {
	if num == 0 {return "0"}
	temp := num
	if temp < 0 {
		temp = -temp
	}
	arr := []rune{}
	for temp != 0 {
		arr = append(arr, rune(temp%7+'0'))
		temp /= 7
	}
	if num < 0 {
		arr = append(arr, '-')
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return string(arr)
}

func Base72(num int) string {
	if num == 0 {
		return "0"
	}
	if num > 0 {
		return dfs(num)
	}
	return "-" + dfs(-num)
}

func dfs(n int) string {
	if n == 0 {
		return ""
	}
	return dfs(n/7) + fmt.Sprintf("%d", n%7)
}