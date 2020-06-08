package main

import "fmt"

var nTest, n, m int
var a [55][55]int

func main() {
	fmt.Scan(&nTest)

	for i := 0; i < nTest; i++ {
		fmt.Scan(&n, &m)
		var row [55]bool
		var col [55]bool

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Scan(&a[i][j])

				if a[i][j] == 1 {
					row[i] = true
					col[j] = true
				}
			}
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
			fmt.Println("Ashish")
		} else {
			fmt.Println("Vivek")
		}

	}
}
