package _021_05

// Decode XORed Permutation
// leetcode: https://leetcode-cn.com/problems/decode-xored-permutation/
// There is an integer array perm that is a permutation of the first n positive integers, where n is always odd.
// It was encoded into another integer array encoded of length n - 1, such that encoded[i] = perm[i] XOR perm[i + 1]. For example, if perm = [1,3,2], then encoded = [2,1].
// Given the encoded array, return the original array perm. It is guaranteed that the answer exists and is unique.

func DecodeXORedPermutation(encoded []int) []int {
	n := len(encoded) + 1
	perm := make([]int, n)
	for i := 1; i <= n; i++ {
		if i % 2 == 1 && i < n {
			perm[0] ^= encoded[i]
		}
		perm[0] ^= i
	}
	for i := 1; i < n; i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}
	return perm
}