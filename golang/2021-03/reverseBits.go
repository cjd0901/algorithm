package _021_03

// Reverse Bits
// leetCode: https://leetcode-cn.com/problems/reverse-bits/
// Reverse bits of a given 32 bits unsigned integer.

func ReverseBits(num uint32) uint32 {
	var res uint32 = 0
	var flag uint32 = 1
	for i := 0; i < 32; i++ {
		res += num & flag
		if i != 31 {
			res <<= 1
			num >>= 1
		}
	}
	return res
}