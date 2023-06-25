package main

import "fmt"

func main() {
	var arr2 [16]int

	for i := range arr2 {
		arr2[i] = i
	}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("array index equals value %d:%d\n", i, arr2[i])
	}
}
