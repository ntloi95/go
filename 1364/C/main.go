package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var a [100000]int
var b [100000]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = -1
	}

	if a[0] == 1 {
		b[0] = 0
	}

	miss := []int{}
	for i := n - 1; i > 0; i-- {
		if a[i] != a[i-1] {
			b[i] = a[i-1]
			for j := a[i] - 1; j > a[i-1]; j-- {
				miss = append(miss, j)
			}
		}
	}

	j := 0
	for i := 0; i < n && j < len(miss); i++ {
		if b[i] == -1 {
			b[i] = miss[len(miss)-1-j]
			j++
		}
	}

	for i := 0; i < n; i++ {
		if b[i] == -1 {
			b[i] = a[n-1] + 1
		}
		fmt.Fprintf(out, "%d ", b[i])
	}
	fmt.Fprintln(out)
	out.Flush()
}
