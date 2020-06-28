package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n)

		res := 0

		x := 0
		y := 0

		for n&1 == 0 {
			n /= 2
			x++
		}

		for n%3 == 0 {
			n /= 3
			y++
		}

		if x > y || n != 1 {
			res = -1
		} else {
			res = 2*y - x
		}
		fmt.Fprintln(out, res)
	}

	out.Flush()
}
