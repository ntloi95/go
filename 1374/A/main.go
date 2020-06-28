package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		var x, y, n int
		fmt.Fscan(in, &x, &y, &n)

		res := n/x*x + y
		if res > n {
			res -= x
		}
		fmt.Fprintln(out, res)
	}

	out.Flush()
}
