package main

import "fmt"

func main() {
	//s1 := make([]byte, 10)
	//s2 := []byte{'g', 'r', 'e', 'a', 't'}
	//s1 = Append(s1, s2)
	buf := []byte{1, 2, 3, 4, 5, 6}
	n := 3

	i := [][]byte{buf[:n], buf[n:]}
	fmt.Println(i)
}
func Append(slice []byte, data []byte) []byte {
	newLength := len(data) + len(slice)
	if cap(slice) < newLength {
		newSlice := make([]byte, cap(slice)+len(data))
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = append(slice, data...)
	return slice
}
