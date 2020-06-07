package main

import "fmt"

func main() {
	var nTest int
	fmt.Scan(&nTest)

	for t := 0; t < nTest; t++ {
		var x, res int64
		fmt.Scan(&x)

		for i := 0; x != 0; i++ {
			if x%2 != 0 {
				res += (1 << (i + 1)) - 1
			}

			x /= 2
		}
		
		fmt.Println(res)
	}
}
