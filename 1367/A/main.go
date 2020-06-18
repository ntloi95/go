package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var s string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)

		for j := 0; j < len(s); j++ {
			if j%2 == 0 || j == len(s)-1 {
				fmt.Fprintf(out, "%c", s[j])
			}
		}

		fmt.Fprintln(out)
	}

	out.Flush()
}
