package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, a, b int
var pos [200001]int
var shift [200001]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		pos[a] = i
	}
	fmt.Fscanln(in)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b)

		s := pos[b] - i
		if s < 0 {
			s += n
		}

		shift[i] = s
	}

	var cnt [200001]int
	res := 0
	for i := 0; i < n; i++ {
		cnt[shift[i]]++
		if cnt[shift[i]] > res {
			res = cnt[shift[i]]
		}
	}

	fmt.Fprint(out, res)
	out.Flush()
}
