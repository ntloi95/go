package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var nTest, n, k int
	var a [200000]int64
	var w [200000]int

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n, &k)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		for i := 0; i < k; i++ {
			fmt.Fscan(in, &w[i])
		}

		sort.Slice(a[:n], func(i, j int) bool {
			return a[i] > a[j]
		})

		sort.Slice(w[:k], func(i, j int) bool {
			return w[i] < w[j]
		})

		res := int64(0)
		j := k - 1
		for i := 0; i < k; i++ {
			res += a[i]
			if w[i] != 1 {
				j += w[i] - 1
				res += a[j]
			} else {
				res += a[i]
			}
		}
		fmt.Fprintln(out, res)
	}

	out.Flush()
}
