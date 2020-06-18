package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, nTest int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n)
		var x int
		even := 0
		odd := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &x)
			if x%2 != j%2 {
				if x%2 == 0 {
					even++
				} else {
					odd++
				}
			}
		}

		if even != odd {
			fmt.Fprintln(out, -1)
			continue
		}

		fmt.Fprintln(out, even)
	}

	out.Flush()
}
