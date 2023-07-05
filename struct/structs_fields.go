package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	//ms := new(struct1)
	//ms.i1 = 10
	//ms.f1 = 15.5
	//ms.str = "Chris"
	// &struct1{} 是一种简写，底层用的也是new()
	ms := &struct1{10, 11.1, "Chris"}

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float32 is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
}
