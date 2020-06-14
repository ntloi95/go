package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n, x, m int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for i := 0; i < nTest; i++ {
		fmt.Fscan(in, &n, &x, &m)

		L, R := x, x
		var l, r int
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &l, &r)

			if (l <= R && R <= r) || (l <= L && L <= r) {
				if l < L {
					L = l
				}

				if r > R {
					R = r
				}
			}
		}

		fmt.Fprintln(out, R-L+1)
	}
	out.Flush()
}
