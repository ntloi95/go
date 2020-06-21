package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(3*(n+1) + 1)
	for i := 0; i < n+1; i++ {
		fmt.Println(i+1, i)
		fmt.Println(i, i)
		fmt.Println(i, i+1)
	}
	fmt.Println(n+1, n+1)
}
