package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)

	fmt.Println(sl_to)
	fmt.Printf("Copied %d element\n", n)

	sl3 := []byte{1, 2, 3}
	//sl3 := []int{1, 2, 3}
	//sl3 = append(sl3, 4, 5, 6)
	bytes := make([]byte, 3)
	AppendByte(sl3, bytes...)
	fmt.Println(sl3)
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:n]
	copy(slice[m:n], data)
	return slice
}
