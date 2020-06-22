package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n int

func gcd(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}

	return x
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n)

		fmt.Fprintln(out, gcd(n/2, n/2*2))
	}

	out.Flush()
}
