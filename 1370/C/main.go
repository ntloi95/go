package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func isPrime(x int) bool {
	sqrtX := int(math.Sqrt(float64(x)))

	for i := 2; i <= sqrtX; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var nTest, n int

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n)

		var res string

		if n == 1 {
			res = "FastestFinger"
		} else if n == 2 {
			res = "Ashishgup"
		} else {
			d := n
			for d&1 == 0 && d > 0 {
				d /= 2
			}

			if (n&(n-1) == 0) || (isPrime(d) && n/d == 2) {
				res = "FastestFinger"
			} else {
				res = "Ashishgup"
			}
		}

		fmt.Fprintln(out, res)
	}

	out.Flush()
}
