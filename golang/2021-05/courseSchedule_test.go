package _021_05

import (
	"fmt"
	"testing"
)

func TestCourseSchedule(t *testing.T) {
	b := CourseSchedule(2, [][]int{
		{1, 0},
	})
	fmt.Println(b)
}