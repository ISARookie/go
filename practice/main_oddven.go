package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if isEven(i) {
			fmt.Printf("%d 是偶数\n", i)
		} else {
			fmt.Printf("%d 是奇数\n", i)
		}
	}
}

func isEven(num int) bool {
	return num%2 == 0
}
