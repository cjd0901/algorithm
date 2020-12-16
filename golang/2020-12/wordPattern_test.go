package golang

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {
	//pattern := WordPattern("abba", "dog dog dog dog")
	//pattern := WordPattern("abba", "dog dog dog dog")
	//pattern := WordPattern("abba", "dog dog dog dog")
	pattern := WordPattern("abba", "dog cat cat dog" +
		"")
	fmt.Println(pattern)
}