package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var nTest, n, k int
var t, a, b [200010]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &n, &k)
	both := []int{}
	alice := []int{}
	bob := []int{}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i], &a[i], &b[i])

		if a[i] == 1 && b[i] == 1 {
			both = append(both, i)
		} else if a[i] == 1 {
			alice = append(alice, i)
		} else if b[i] == 1 {
			bob = append(bob, i)
		}
	}

	sort.Slice(both, func(i, j int) bool {
		return t[both[i]] < t[both[j]]
	})
	sort.Slice(alice, func(i, j int) bool {
		return t[alice[i]] < t[alice[j]]
	})
	sort.Slice(bob, func(i, j int) bool {
		return t[bob[i]] < t[bob[j]]
	})

	res := 0
	x := 0
	y := 0
	z := 0
	for i := 0; i < k; i++ {
		if x >= len(both) && (y >= len(alice) || z >= len(bob)) {
			fmt.Fprintln(out, -1)
			out.Flush()
			return
		}

		if x >= len(both) {
			res += t[alice[y]] + t[bob[z]]
			y++
			z++
		} else if y >= len(alice) || z >= len(bob) {
			res += t[both[x]]
			x++
		} else {
			if t[both[x]] <= t[alice[y]]+t[bob[z]] {
				res += t[both[x]]
				x++
			} else {
				res += t[alice[y]] + t[bob[z]]
				y++
				z++
			}
		}
	}

	fmt.Fprintln(out, res)
	out.Flush()
}
