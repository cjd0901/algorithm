package _021_03

// Minimum Domino Rotations For Equal Row
// leetCode: https://leetcode-cn.com/problems/minimum-domino-rotations-for-equal-row/
// In a row of dominoes, A[i] and B[i] represent the top and bottom halves of the ith domino. 
// (A domino is a tile with two numbers from 1 to 6 - one on each half of the tile.)
// We may rotate the ith domino, so that A[i] and B[i] swap values.
// Return the minimum number of rotations so that all the values in A are the same, or all the values in B are the same.
// If it cannot be done, return -1.

func MinDominoRotations(A, B []int) int {
	nums := [7]int{}
	a := [7]int{}
	b := [7]int{}
	for i := range A {
		if A[i] == B[i] {
			nums[A[i]]++
		}else {
			nums[A[i]]++
			nums[B[i]]++
			a[A[i]]++
			b[B[i]]++
		}
	}
	for i, n := range nums {
		if n == len(A) {
			if a[i] > b[i] {
				return b[i]
			}else {
				return a[i]
			}
		}
	}
	return -1
}

//func MinDominoRotations(A, B []int) int {
//	sort.Ints(A)
//	sort.Ints(B)
//	am := make(map[int]int)
//	bm := make(map[int]int)
//	for i := range A {
//		am[A[i]] ++
//		bm[B[i]] ++
//	}
//	l := len(A)
//	for k := range am {
//		sum := am[k] + bm[k]
//		if sum >= l {
//			if am[k] > bm[k] {
//				return l - am[k]
//			}else {
//				return l - bm[k]
//			}
//		}
//	}
//	return -1
//}
