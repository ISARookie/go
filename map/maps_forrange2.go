package main

import "fmt"

func main() {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		// value 为map [1] 为key
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	items2 := make([]map[int]int, 5)
	// item 为value的copy不能进行修改同数组
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}
