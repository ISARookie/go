package main

import "fmt"

func main() {
	fibseq := fibonacci2(6)
	fmt.Println(fibseq)
}

func fibonacci2(n int) []int {
	seq := make([]int, n)
	if n < 1 {
		return seq
	} else if n == 1 {
		seq[0] = 1
		return seq
	} else if n > 1 {
		seq[0] = 1
		seq[1] = 1
	}

	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq
}
