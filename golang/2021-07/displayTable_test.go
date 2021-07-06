package _021_07

import (
	"fmt"
	"testing"
)

func TestDisplayTable(t *testing.T) {
	ans := DisplayTable([][]string{
		{"David","3","Ceviche"},
		{"Corina","10","Beef Burrito"},
		{"David","3","Fried Chicken"},
		{"Carla","5","Water"},
		{"Carla","5","Ceviche"},
		{"Rous","3","Ceviche"},
	})
	fmt.Println(ans)
}