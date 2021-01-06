package _021_01

import (
	"fmt"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	values := []float64{
		2.0,
		3.0,
	}
	queries := [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}
	//equations := [][]string{
	//	{"a", "e"},
	//	{"b", "e"},
	//}
	//values := []float64{
	//	4.0,
	//	3.0,
	//}
	//queries := [][]string{
	//	{"a", "b"},
	//	{"e", "e"},
	//	{"x", "x"},
	//}
	//equations := [][]string{
	//	{"a", "b"},
	//	{"e", "f"},
	//	{"b", "e"},
	//}
	//values := []float64{
	//	3.4,
	//	1.4,
	//	2.3,
	//}
	//queries := [][]string{
	//	{"b", "a"},
	//	{"a", "f"},
	//	{"f", "f"},
	//	{"e", "e"},
	//	{"c", "c"},
	//	{"a", "c"},
	//	{"f", "e"},
	//}
	ans := CalcEquation(equations, values, queries)
	fmt.Println(ans)
}