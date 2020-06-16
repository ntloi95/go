package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n int
var a [100000]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		res := []int{a[0]}
		for i := 1; i < n-1; i++ {
			if (a[i-1] < a[i]) != (a[i] < a[i+1]) {
				res = append(res, a[i])
			}
		}
		res = append(res, a[n-1])

		fmt.Fprintln(out, len(res))
		for _, v := range res {
			fmt.Fprintf(out, "%d ", v)
		}
		fmt.Fprintln(out)
	}
	out.Flush()
}
