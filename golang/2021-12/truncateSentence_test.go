package _021_12

import (
	"fmt"
	"testing"
)

func TestTruncateSentence(t *testing.T) {
	ans := TruncateSentence("Hello how are you Contestant", 4)
	fmt.Println(ans)
}