package main

import "fmt"

// make /** 创建数组切片的时候会在一个数组的基础上进行切片在没有数组声明的情况下切片的声明无法变更？
// cap /** 明确cap的值为[start:]
// make /** 数组切片为数组的一部分的引用，可以在数组过大的情况下切取合适的位置数组传输
func main() {
	//s := make([]byte, 5)
	//s = s[2:4]
	//fmt.Println(len(s))
	//fmt.Println(cap(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	s2[1] = 't'
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}
}
