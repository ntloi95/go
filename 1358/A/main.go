package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n, m int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &m, &n)

		res := (n*m + 1) / 2

		fmt.Fprintln(out, res)
	}

	out.Flush()
}
