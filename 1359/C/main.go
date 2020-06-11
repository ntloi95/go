package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var s int
	var i int
	fmt.Fscanf(in, "%v\n%c", &i, &s)

	fmt.Println(s, byte('.'))
	out.Flush()
}
