package _021_05

import (
	"fmt"
	"testing"
)

func TestEmployeeImportance(t *testing.T) {
	employees := []*Employee{
		{1, 5, []int{2,3}},
		{2, 3, []int{}},
		{3, 3, []int{}},
	}
	importance := EmployeeImportance(employees, 1)
	fmt.Println(importance)
}