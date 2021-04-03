package _021_04

import (
	"fmt"
	"testing"
)

func TestVolumeOfHistogramLCCI(t *testing.T) {
	volume := VolumeOfHistogramLCCI([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	//volume := VolumeOfHistogramLCCI([]int{2, 1, 0, 2})
	fmt.Println(volume)
}