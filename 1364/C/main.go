package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var a [100000]int
var b [100000]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = -1
	}

	miss := []int{}
	for i := n - 1; i > 0; i-- {
		if a[i] != a[i-1] {
			b[i] = a[i-1]
			from := a[i-1] + 1
			to := a[i]

			for j := from; j < to && i > 0; j++ {

				if i > 1 && a[i-1] != a[i-2] {
					// fmt.Println(a[i-1], a[i-2])
					// fmt.Println(from, to)
					for k := to - 1; k >= j; k-- {
						// fmt.Println(k)
						miss = append(miss, k)
					}
					break
				}

				b[i-1] = j
				i--
			}
		}
	}

	// for _, x := range miss {
	// 	// fmt.Println(x)
	// }

	j := 0
	for i := 0; i < n && j < len(miss); i++ {
		if b[i] == -1 {
			if i == 0 && a[i] == 1 {
				b[i] = 0
			} else {
				b[i] = miss[len(miss)-1-j]
				j++
			}
		}
	}

	if j != len(miss) {
		fmt.Print(-1)
		return
	}

	val := 0
	for i := 0; i < n; i++ {
		if b[i] == -1 {
			if a[i] == 0 {
				b[i] = 100001
			} else {
				b[i] = val
			}
		} else {
			val = b[i]
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", b[i])
	}
	fmt.Fprintln(out)
	out.Flush()
}
