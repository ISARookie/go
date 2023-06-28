package main

import "fmt"

func main() {
	//var mapLit map[string]int
	//var mapAssigned map[string]int
	//
	//mapLit = map[string]int{"one": 1, "two": 2}
	//
	//mapCreated := make(map[string]float32)
	//mapAssigned = mapLit
	//
	//mapCreated["key1"] = 4.5
	//mapCreated["key2"] = 3.14159
	//mapCreated["two"] = 3

	m := map[int]func() int{
		1: func() int {
			return 10
		},
	}
	fmt.Println(m)
	fmt.Println(m[1]())
}
