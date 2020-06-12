package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var nTest int
var h, c, t int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &nTest)

	for i := 0; i < nTest; i++ {
		fmt.Fscan(in, &h, &c, &t)

		if t <= (h+c)/2 {
			fmt.Fprintln(out, 2)
			continue
		}

		k := ((h - t) / (2*t - h - c))

		v1 := (k+1)*h + k*c
		m1 := 2*k + 1
		v2 := (k+2)*h + (k+1)*c
		m2 := 2*(k+1) + 1

		mm := m1 * m2
		V1 := v1 * m2
		V2 := v2 * m1
		T := t * mm

		if math.Abs(float64(V1-T)) <= math.Abs(float64(V2-T)) {
			fmt.Fprintln(out, k*2+1)
		} else {
			fmt.Fprintln(out, (k+1)*2+1)
		}
	}
	out.Flush()
}
