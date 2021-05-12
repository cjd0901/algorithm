package _021_05

import (
	"fmt"
	"testing"
)

func TestBulbSwitcherIV(t *testing.T) {
	count := BulbSwitcherIV("10111")
	fmt.Println(count)
}