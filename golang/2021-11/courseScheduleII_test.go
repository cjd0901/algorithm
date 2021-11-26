package _021_11

import (
	"fmt"
	"testing"
)

func TestCourseScheduleIIDFS(t *testing.T) {
	ans := CourseScheduleIIDFS(2, [][]int{
		{1, 0},
	})
	fmt.Println(ans)
}