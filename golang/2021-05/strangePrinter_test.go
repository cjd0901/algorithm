package _021_05

import (
	"fmt"
	"testing"
)

func TestStrangePrinter(t *testing.T) {
	counts := StrangePrinter("aba")
	fmt.Println(counts)
}