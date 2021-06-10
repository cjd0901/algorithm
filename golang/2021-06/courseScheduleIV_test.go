package _021_06

import (
	"fmt"
	"testing"
)

func TestCourseScheduleIV(t *testing.T) {
	ans := CourseScheduleIV(5, [][]int{
		{1, 0},
		{2, 0},
	}, [][]int{
		{0, 1},
		{2, 0},
	})
	fmt.Println(ans)
}