package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest int
var a, b, c, x1, x2 int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &a, &b, &c)

		if a < c {
			x1 = 1
		} else {
			x1 = -1
		}

		if a*b <= c {
			x2 = -1
		} else {
			x2 = b
		}
		fmt.Fprintln(out, x1, x2)
	}

	out.Flush()
}
