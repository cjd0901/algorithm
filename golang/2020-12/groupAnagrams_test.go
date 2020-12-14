package golang

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	anagrams := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(anagrams)
}