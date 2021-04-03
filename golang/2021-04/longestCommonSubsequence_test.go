package _021_04

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	ans := LongestCommonSubsequence("abcde", "dabef")
	fmt.Println(ans)
}