package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	a := [200000]int{}
	bit := [20][]bool{}
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < n; j++ {
			b := (a[j] & (1 << i)) != 0
			bit[i] = append(bit[i], b)
		}

		sort.Slice(bit[i], func(x, y int) bool {
			return bit[i][x] && !bit[i][y]
		})
	}

	sum := int64(0)
	for i := 0; i < n; i++ {
		val := int64(0)
		for j := 0; j < 20; j++ {
			if bit[j][i] {
				val += 1 << j
			}
		}

		sum += val * val
	}

	fmt.Fprint(out, sum)
	out.Flush()
}
