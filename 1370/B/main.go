package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n int

type pair struct {
	val int
	id  int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n)
		odd := []int{}
		even := []int{}
		var x int
		for i := 0; i < n*2; i++ {
			fmt.Fscan(in, &x)

			if x&1 == 0 {
				even = append(even, i+1)
			} else {
				odd = append(odd, i+1)
			}
		}

		nOdd := len(odd) / 2 * 2
		nEven := 2*n - 2 - nOdd
		if nEven < 0 {
			nEven = 0
			nOdd -= 2
		}
		// fmt.Println("AHIHI", nOdd, nEven)
		for i := 0; i < nOdd/2; i++ {
			fmt.Fprintln(out, odd[i*2], odd[i*2+1])
		}
		for i := 0; i < nEven/2; i++ {
			fmt.Fprintln(out, even[i*2], even[i*2+1])
		}
	}

	out.Flush()
}
