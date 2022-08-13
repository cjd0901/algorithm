package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	j := (n + 1) / 2
	for i := 1; i < j; i++ {
		sum := i
		for k := i + 1; k <= j; k++ {
			sum += k
			if sum == n {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}
