package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n int
var s string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n, &s)

		cnt := 0
		for i := 0; i < n; i++ {
			if s[i] == ')' && cnt > 0 {
				cnt--
			}

			if s[i] == '(' {
				cnt++
			}
		}
		fmt.Fprintln(out, cnt)
	}

	out.Flush()
}
