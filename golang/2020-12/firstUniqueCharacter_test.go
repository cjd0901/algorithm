package golang

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	index := FirstUniqChar("loveleetcode")
	//index := FirstUniqChar("leetcode")
	fmt.Println(index)
}