package main

import "fmt"

type C struct {
	f1 float32
	int
	string
}

func main() {
	c := C{2.2, 1, "11"}
	fmt.Println(c)
}
