package main

import (
	. "example.com/m/v2/chapter/chapter9/pack1"
	"fmt"
)

func main() {
	var test1 string
	test1 = ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", Pack1Int)
	fmt.Printf("Float from package1: %f\n", Pack1Float)
}
