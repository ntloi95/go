package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n, m int
var a [30][30]int

func level(x, y int) int {
	d1 := x + y
	d2 := n - 1 - x + m - 1 - y

	if d1 < d2 {
		return d1
	}

	return d2
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n, &m)
		res := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &a[i][j])
			}
		}

		for l := 0; l < (n+m-1)/2; l++ {
			zero := 0
			one := 0
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					lv := level(i, j)

					if lv == l {
						if a[i][j] == 1 {
							one++
						} else {
							zero++
						}
					}
				}
			}
			// fmt.Println(l, zero, one)
			if zero < one {
				res += zero
			} else {
				res += one
			}
		}

		fmt.Fprintln(out, res)
	}
	out.Flush()
}
