package _021_05

// XOR Queries of a Subarray
// leetcode: https://leetcode-cn.com/problems/xor-queries-of-a-subarray/
// Given the array arr of positive integers and the array queries where queries[i] = [Li, Ri], for each query i compute the XOR of elements from Li to Ri (that is, arr[Li] xor arr[Li+1] xor ... xor arr[Ri] ).
// Return an array containing the result for the given queries

func XORQueriesOfASubarray(arr []int, queries [][]int) []int {
	qs := make([]int, len(queries))
	xors := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}
	for i, q := range queries {
		qs[i] = xors[q[0]] ^ xors[q[1]+1]
	}
	return qs
}