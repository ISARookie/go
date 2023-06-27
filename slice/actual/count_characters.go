package main

import (
	"os"
	"regexp"
)

func main() {
	//str1 := "123456789"
	//fmt.Println(utf8.RuneCountInString(str1))
	//
	//s := str1[0:3]
	//fmt.Println(s)
	//var b = []byte{'1', '2', '3', '9'}
	//var a = []byte{'2'}
	//compare := Compare(a, b)
	//fmt.Println(compare)
	// 7.6.8

}

var digitRegexp = regexp.MustCompile("[0-9]+")

// FindDigits /** 避免局部占用导致整个文件的内存得不到释放
func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	//return digitRegexp.Find(b)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// Compare /** a []byte,b []byte
func Compare(a, b []byte) int {
	// 判断交集部分
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return 0
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}
