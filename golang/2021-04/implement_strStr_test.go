package _021_04

import (
	"fmt"
	"testing"
)

func TestImplement_strStr(t *testing.T) {
	index := Implement_strStr("mississippi", "issip")
	fmt.Println(index)
}