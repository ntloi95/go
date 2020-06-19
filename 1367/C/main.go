package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, k, nTest int
var s string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n, &k)
		fmt.Fscan(in, &s)
		res := 0

		dis := k
		for i := 0; i < n+k; i++ {
			if i >= n || s[i] == '0' {
				dis++
			} else {
				dis = 0
			}

			if dis == 2*k+1 {
				res++
				dis = k
			}
		}
		fmt.Fprintln(out, res)
	}

	out.Flush()
}
