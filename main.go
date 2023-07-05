package main

import (
	structPack "example.com/m/v2/struct/pack"
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() {
	go backend()
}

func backend() {

}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" type %T and has value %v \n", g, g)
	}
}

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mil = 10
	struct1.Mfl = 10.

	fmt.Printf("Mil = %d\n", struct1.Mil)
	fmt.Printf("Mfl = %f\n", struct1.Mfl)

	// 闭包
	//f()
	//var goos = os.Getenv("GOOS")
	//fmt.Printf("The operating system is: %s\n", goos)
	//path := os.Getenv("PATH")
	//fmt.Printf("Path is %s\n", path)

	//var a int
	//Func1()
	//fmt.Println(a)
	//const Ln2 = 0.693147180559945309417232121458 /
	//	176568075500134360255254120680009
	//const Log2E = 1 / Ln2
	//const Billion = 1e9
	//const hardEight = (1 << 100) >> 97
	//fmt.Println(Log2E)

	//var n int64 = 2
	//fmt.Println(n)
	//
	//var (
	//	HOME   = os.Getenv("HOME")
	//	USER   = os.Getenv("USER")
	//	GOROOT = os.Getenv("GOROOT")
	//)
	//
	//fmt.Printf("HOME %s,USER%s,GOROOT%s\n", HOME, USER, GOROOT)

}

func (t T) Method1() {

}

func Func1() {

}
