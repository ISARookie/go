package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 7.12
	//str := []string{"a", "b", "c", "d", "e", "f"}
	//var index = 3
	//arr1, arr2 := substring(str, index)
	//fmt.Printf("substring index %d, early part %s :latter part %s\n", index, arr1, arr2)
	// 7.13
	//str := "abcdefg"
	//str1 := str[len(str)/2:] + str[:len(str)/2]
	//fmt.Println(str1)
	// 7.14
	//s := reverseString("Google")
	//fmt.Println(s)
	// 7.15
	//str := []string{"a", "b", "b", "c", "c", "e", "f"}
	//fmt.Println(difArr(str))
	// 7.16
	ints := []int{1, 5, 3, 4, 5, 6, 7, 8}
	//bubbleSort(ints)
	fold := func(data int) int {
		return data * 10
	}
	tenfold(fold, ints)
	fmt.Println(ints)
}
func tenfold(a func(int) int, ints []int) {
	for index, value := range ints {
		ints[index] = a(value)
	}
}

func bubbleSort(slice []int) {
	length := len(slice)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func difArr(arr []string) []string {
	result := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if i != 0 && arr[i] == arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result
}

func reverseString(s string) string {
	rune := []rune(s)
	length := len(rune)
	for i := 0; i < length/2; i++ {
		rune[i], rune[length-i-1] = rune[length-i-1], rune[i]
	}
	return string(rune)
}

func reverse(data string) string {
	arr := []byte(data)
	buffer := bytes.Buffer{}
	for i := len(arr) - 1; i >= 0; i-- {
		buffer.WriteByte(arr[i])
	}
	return buffer.String()
}
func substring(arr []string, index int) (arr1 []string, arr2 []string) {
	return arr[:index], arr[index:]
}
