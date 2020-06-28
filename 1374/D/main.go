package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var nTest, n int
var k int64
var a [200001]int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &nTest)
	for t := 0; t < nTest; t++ {
		fmt.Fscan(in, &n, &k)
		cnt := map[int64]int{}
		mx := int64(0)
		vmx := int64(0)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
			a[i] %= int64(k)
			if a[i] != 0 {
				a[i] = int64(k) - a[i]
			}

		}

		sort.Slice(a[:n], func(i, j int) bool {
			return a[i] < a[j]
		})

		// fmt.Println(a[:n])
		for i := 0; i < n; i++ {
			if a[i] != 0 {
				cnt[a[i]]++
			}

			if cnt[a[i]] >= int(mx) {
				mx = int64(cnt[a[i]])
				vmx = a[i]
			}
		}

		// fmt.Println(mx, vmx)
		var res int64
		if mx == 0 {
			res = 0
		} else {
			res = k*(mx-1) + vmx + 1
		}
		fmt.Fprintln(out, res)

	}

	out.Flush()
}
