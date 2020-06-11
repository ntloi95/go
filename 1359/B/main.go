package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n, m, x, y int
var a [100][1000]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n, &m, &x, &y)
		fmt.Fscanln(in)
		cntDot := 0

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscanf(in, "%c", &a[i][j])
				if a[i][j] == '.' {
					cntDot++
				}
			}
			fmt.Fscanln(in)
		}

		if 2*x <= y {
			fmt.Fprintln(out, x*cntDot)
			continue
		}

		cnt1x2 := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if a[i][j] == '.' && j < m-1 && a[i][j+1] == '.' {
					a[i][j] = '*'
					a[i][j+1] = '*'
					cnt1x2++
				}
			}
		}
		fmt.Fprintln(out, x*(cntDot-cnt1x2*2)+y*cnt1x2)
	}
	out.Flush()
}
