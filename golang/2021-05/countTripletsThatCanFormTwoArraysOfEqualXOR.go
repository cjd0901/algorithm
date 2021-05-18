package _021_05

// Count Triplets That Can From Two Arrays of Equal XOR
// leetcode: https://leetcode-cn.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
// Given an array of integers arr.
// We want to select three indices i, j and k where (0 <= i < j <= k < arr.length).
// Let's define a and b as follows:
// a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
// b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
// Note that ^ denotes the bitwise-xor operation.
// Return the number of triplets (i, j and k) Where a == b.

func CountTripletsThatCanFormTwoArraysOfEqualXOR(arr []int) int {
	n := len(arr)
	s := make([]int, n+1)
	for i := range arr {
		s[i+1] = s[i] ^ arr[i]
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				if s[i] == s[k+1] {
					ans++
				}
			}
		}
	}
	return ans
}