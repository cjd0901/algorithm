package _021_01

// Binary Prefix Divisible By 5
// leetCode: https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/
// Given an array A of 0s and 1s, consider N_i: the i-th subarray from A[0] to A[i] interpreted as a binary number (from most-significant-bit to least-significant-bit.)
// Return a list of booleans answer, where answer[i] is true if and only if N_i is divisible by 5.

func PrefixesDivBy5(A []int) []bool {
	ans := make([]bool, len(A))
	x := 0
	for i, v := range A {
		x = (x<<1 | v) % 5
		ans[i] = x == 0
	}
	return ans
}