package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, a, b int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for i := 0; i < nTest; i++ {
		fmt.Fscan(in, &a, &b)

		if a < b {
			a, b = b, a
		}

		if a >= 2*b {
			fmt.Fprintln(out, b)
			continue
		}

		t := 0
		if a != b {
			t = a - b
		}

		res := t + (a-2*t)/3*2
		if (a-2*t)%3 == 2 {
			res++
		}
		fmt.Fprintln(out, res)
	}
	out.Flush()
}
