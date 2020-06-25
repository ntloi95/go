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
		first1ID := -1
		last0ID := -1

		for i := 0; i < n; i++ {
			if s[i] == '1' && first1ID == -1 {
				first1ID = i
			}

			if s[i] == '0' {
				last0ID = i
			}
		}

		if last0ID == -1 {
			for i := 0; i < n; i++ {
				fmt.Fprint(out, 1)
			}
		} else if first1ID == -1 {
			for i := 0; i < n; i++ {
				fmt.Fprint(out, 0)
			}
		} else {
			for i := 0; i < first1ID; i++ {
				fmt.Fprint(out, 0)
			}
			if first1ID < last0ID {
				fmt.Fprint(out, 0)
			}
			for i := n - 1; i > last0ID; i-- {
				fmt.Fprint(out, 1)
			}
		}

		fmt.Fprintln(out)
	}

	out.Flush()
}
