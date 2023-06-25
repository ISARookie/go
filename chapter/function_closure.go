package main

import (
	"fmt"
	"strings"
)

func main() {
	//var f = Adder()
	//fmt.Print(f(1), " - ")
	//fmt.Print(f(20), " - ")
	//fmt.Print(f(300))

	//var g int
	//go func(i int) {
	//	s := 0
	//	for j := 0; j < i; j++ {
	//		s += j
	//		g = s
	//	}
	//}(1000)
	//fmt.Print()
	addBmp := MakeAddSuffix(".bmp")
	fmt.Println(addBmp("file"))
}

// MakeAddSuffix /** 返回值是一个函数，通过入参自定义返回形式
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
