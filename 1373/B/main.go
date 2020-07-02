package main

import (
	"bufio"
	"fmt"
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

		cnt1 := 0
		cnt0 := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}

		min := cnt0
		if cnt1 < cnt0 {
			min = cnt1
		}

		var res string

		if min&1 == 0 {
			res = "NET"
		} else {
			res = "DA"
		}

		fmt.Fprintln(out, res)
	}

	out.Flush()
}
