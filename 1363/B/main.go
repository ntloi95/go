package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var nTest int
var s string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &s)
		l := len(s)
		cnt0 := make([]int, l+1)
		cnt1 := make([]int, l+1)

		for i := 1; i <= l; i++ {
			if s[i-1] == '0' {
				cnt0[i] = cnt0[i-1] + 1
				cnt1[i] = cnt1[i-1]
			} else {
				cnt1[i] = cnt1[i-1] + 1
				cnt0[i] = cnt0[i-1]
			}
		}

		res := math.MaxInt32
		for i := 0; i < l; i++ {
			res01 := cnt0[i] + cnt1[l] - cnt1[i+1]
			res10 := cnt1[i] + cnt0[l] - cnt0[i+1]

			if res > res01 {
				res = res01
			}
			if res > res10 {
				res = res10
			}
		}

		fmt.Fprintln(out, res)
	}

	out.Flush()
}
