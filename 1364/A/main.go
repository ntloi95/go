package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n, x int
var a [100000]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for i := 0; i < nTest; i++ {
		fmt.Fscan(in, &n, &x)

		sum := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
			sum += a[i]
		}

		sum1 := sum
		i1 := 0
		for sum1%x == 0 && i1 < n {
			sum1 -= a[i1]
			i1++
		}

		sum2 := sum
		i2 := n - 1
		for sum2%x == 0 && i2 >= 0 {
			sum2 -= a[i2]
			i2--
		}

		len1 := 0
		len2 := 0
		if i1 < n {
			len1 = n - i1
		}

		if i2 >= 0 {
			len2 = i2 + 1
		}

		if len1 == 0 && len2 == 0 {
			fmt.Fprintln(out, -1)
		} else if len1 < len2 {
			fmt.Fprintln(out, len2)
		} else {
			fmt.Fprintln(out, len1)
		}
	}
	out.Flush()
}
