package main

import "fmt"

func main() {
	var arr [3]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
		fmt.Println(&arr)
	}
}
