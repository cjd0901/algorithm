package _021_05

import (
	"fmt"
	"testing"
)

func TestDeleteAndEarn(t *testing.T) {
	earn := DeleteAndEarn([]int{3, 4, 2})
	fmt.Println(earn)
}