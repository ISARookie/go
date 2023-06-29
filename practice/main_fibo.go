package main

import (
	"example.com/m/v2/practice/fibo"
	"example.com/m/v2/practice/fibona"
	"fmt"
)

func main() {
	fibona.Fbb(12)

	// 使用斐波那契包中的全局变量进行操作
	fmt.Println("Sum of last Fibonacci value and 5:", fibo.LastValue+fibo.Operation(5))
	fmt.Println("Product of last Fibonacci value and 3:", fibo.LastValue*fibo.Operation(3))
}
