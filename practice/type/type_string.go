package main

import "fmt"

type T struct {
	a int
	b float32
	c string
}

func (t T) String() {
	fmt.Printf("%v\n", t)
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	t.String()
}
