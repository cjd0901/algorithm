package golang

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	letters := RemoveDuplicateLetters("cbacdcbc")
	fmt.Println(letters)
}