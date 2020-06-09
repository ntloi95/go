package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n int
var a [555]int
var b [555]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for i := 0; i < nTest; i++ {
		fmt.Fscan(in, &n)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		fmt.Fscanln(in)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &b[i])
		}
		fmt.Fscanln(in)

		nonDecrease := true
		oneType := true

		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				nonDecrease = false
			}

			if b[i] != b[i+1] {
				oneType = false
			}

		}

		if oneType && !nonDecrease {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
	}
	out.Flush()
}
