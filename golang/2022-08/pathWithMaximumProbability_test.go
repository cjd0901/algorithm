package main

import "testing"

func TestPathWithMaximumProbability(t *testing.T) {
	edges := [][]int{
		{0, 1},
	}
	succProb := []float64{0.5}
	t.Log(PathWithMaximumProbability(3, edges, succProb, 0, 2))
}
