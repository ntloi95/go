package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n, m, k int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for i := 0; i < nTest; i++ {
		fmt.Fscan(in, &n, &m, &k)

		maxJokers := n / k
		if maxJokers > m {
			maxJokers = m
		}

		secondMaxJoker := (m - maxJokers + k - 2) / (k - 1)

		fmt.Fprintln(out, maxJokers-secondMaxJoker)
	}
	out.Flush()
}
