package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n, m int
var a [55][55]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for i := 0; i < nTest; i++ {
		fmt.Fscan(in, &n, &m)
		var row [55]bool
		var col [55]bool

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &a[i][j])

				if a[i][j] == 1 {
					row[i] = true
					col[j] = true
				}
			}
			fmt.Fscanln(in)
		}

		nRow := 0
		for i := 0; i < n; i++ {
			if !row[i] {
				nRow++
			}
		}

		nCol := 0
		for i := 0; i < m; i++ {
			if !col[i] {
				nCol++
			}
		}

		var step int
		if nRow < nCol {
			step = nRow
		} else {
			step = nCol
		}

		if step&1 == 1 {
			fmt.Fprintln(out, "Ashish")
		} else {
			fmt.Fprintln(out, "Vivek")
		}
	}
	out.Flush()
}
