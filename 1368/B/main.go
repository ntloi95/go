package main

import "fmt"

func main() {
	var k int64
	s := "codeforces"
	cnt := [10]int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	fmt.Scan(&k)

	sum := int64(1)
	for i := 0; sum < k; i++ {
		i %= 10
		cnt[i]++
		sum = sum / (cnt[i] - 1) * cnt[i]
	}

	for i := 0; i < 10; i++ {
		for j := 0; int64(j) < cnt[i]; j++ {
			fmt.Printf("%c", s[i])
		}
	}
}
