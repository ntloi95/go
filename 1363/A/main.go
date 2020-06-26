package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n, x, a int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n, &x)
		odd := 0
		even := 0

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a)
			if a&1 == 0 {
				even++
			} else {
				odd++
			}
		}

		odd--
		if odd < 0 {
			fmt.Fprintln(out, "No")
			continue
		}

		if x&1 == 0 {
			even--

			if even < 0 {
				fmt.Fprintln(out, "No")
				continue
			}
		}

		if (x-1)/2 > (odd/2 + even/2) {
			fmt.Fprintln(out, "No")
			continue
		}

		fmt.Fprintln(out, "Yes")
	}

	out.Flush()
}
