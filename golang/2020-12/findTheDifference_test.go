package golang

import (
	"fmt"
	"testing"
)

func TestFindTheDifference(t *testing.T) {
	difference := FindTheDifference4("adcb", "abcde")
	fmt.Println(difference)
}