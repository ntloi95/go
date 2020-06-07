package main

import "fmt"

func main() {
	var nTest int
	fmt.Scan(&nTest)

	for t := 0; t < nTest; t++ {
		var n int
		fmt.Scan(&n)
		var a [1025]int

		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		var cnt [1025]int
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				cnt[a[i]^a[j]]++
			}
		}

		hasSolution := false
		for i := 1; i < 1025; i++ {
			if cnt[i] == n {
				fmt.Println(i)
				hasSolution = true
				break
			}
		}

		if !hasSolution {
			fmt.Println(-1)
		}
	}
}
