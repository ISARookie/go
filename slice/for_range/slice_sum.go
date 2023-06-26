package main

import "fmt"

// 变长参数...type /** sum(seq...) 展开数组或者切片的数据作为入参
func main() {
	seq := []int{1, 2, 3, 4, 5, 6, 7}
	total := sum(seq...)
	fmt.Println(total)
}

func sum(data ...int) int {
	var total int
	for _, value := range data {
		total += value
	}
	return total
}
