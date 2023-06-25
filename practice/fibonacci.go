package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result := fibonacci(i)
		fmt.Printf("fibonacci(%d) is : %d\n", i, result)
		end := time.Now()
		delta := end.Sub(start)
		fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	}
}

// LIM /** 数组长度
const LIM = 42

var fibs [LIM]uint64

// fibonacci /** n = (n-1) + (n-2)
// 利用闭包变量map存值来判断是否已经计算过该值
func fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		fmt.Printf("fibonacci(%d) result : %d \n", n, res)
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
