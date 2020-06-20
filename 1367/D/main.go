package main

import (
	"bufio"
	"fmt"
	"os"
)

var nTest, n int
var a [55]int
var s string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &s, &n)
		var res [55]int
		cnt := map[int]int{}
		for i := 0; i < len(s); i++ {
			cnt[int(s[i])]++
		}

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		fill := 0
		char := int('z')
		for fill != n {
			zero := []int{}
			for i := 0; i < n; i++ {
				if a[i] == 0 && res[i] == 0 {
					zero = append(zero, i)
				}
			}
			for cnt[char] < len(zero) {
				char--
			}
			cnt[char] -= len(zero)
			fill += len(zero)
			for j := 0; j < len(zero); j++ {
				res[zero[j]] = char
				for k := 0; k < n; k++ {
					if a[k] != 0 {
						if zero[j] > k {
							a[k] -= zero[j] - k
						} else {
							a[k] -= k - zero[j]
						}
					}
				}
			}
			char--
		}

		for i := 0; i < n; i++ {
			fmt.Fprintf(out, "%c", res[i])
		}
		fmt.Fprintln(out)
	}

	out.Flush()
}
