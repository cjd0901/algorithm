package _021_01

import (
	"fmt"
	"testing"
)

func TestLargeGroupPositions(t *testing.T) {
	//positions := LargeGroupPositions("abbxxxxzzy")
	positions := LargeGroupPositions("abbxxxxzzy")
	//positions := LargeGroupPositions("abc")
	//positions := LargeGroupPositions("abcdddeeeeaabbbcd")
	//positions := LargeGroupPositions("aba")
	//positions := LargeGroupPositions("aaa")
	fmt.Println(positions)
}