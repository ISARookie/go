package main

import "fmt"

// for-range /** value 为数组值的拷贝修改不影响原始值
func main() {
	seq := []byte{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}
	n := 5
	selices := [][]byte{seq[:n], seq[n:]}

	for row := range selices {
		for column := range selices[row] {
			selices[row][column] *= 2
		}
	}

	for row := range selices {
		for column := range selices[row] {
			fmt.Println(selices[row][column])
		}
	}
}
