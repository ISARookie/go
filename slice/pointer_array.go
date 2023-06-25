package main

import "fmt"

// main /** *p 获取地址对应的实际值 &p 获取p的内存地址
func main() {
	var ar [3]int
	f(ar)
	fp(&ar)
	fmt.Println(&ar)
}

func fp(a *[3]int) {
	fmt.Println(a)
}

func f(a [3]int) {
	fmt.Println(a)
}
