package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, a, b, n int

func count(x, y, z int) int {
	cnt := 0
	for x <= z && y <= z {
		x += y
		cnt++
		if x > z {
			break
		}
		y += x
		cnt++
	}

	return cnt
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &a, &b, &n)

		if a > b {
			a, b = b, a
		}
		cnt2 := count(a, b, n)

		fmt.Fprintln(out, cnt2)
	}

	out.Flush()
}
