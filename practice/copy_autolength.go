package main

import "fmt"

func main() {
	// 7.9
	//slice := make([]int, 4)
	//var factor = 4
	//newSlice := make([]int, len(slice)*factor)
	//copy(slice, newSlice)
	//slice = newSlice
	//for i := 0; i < len(slice); i++ {
	//	fmt.Println(slice[i])
	//}
	//fmt.Printf("Slice length is :%d\n", len(slice))
	// 7.10
	//s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fn := func(num int) bool {
	//	return num%2 == 0
	//}
	//filtered := Filter(s, fn)
	//fmt.Println(filtered)
	// 7.11
	//slice1 := []string{"a", "b", "c", "d"}
	//slice2 := []string{"x", "y", "z"}
	//
	//// 在 slice1 的索引 2 处插入 slice2
	//result := InsertStringSlice(slice1, 2, slice2)
	//
	//fmt.Println("Original Slice:", slice1)
	//fmt.Println("Insert Slice:", slice2)
	//fmt.Println("Result:", result)

	slice1 := []string{"a", "b", "c", "d"}
	slice := RemoveStringSlice(slice1, 1, 1)
	fmt.Println(slice)
}

func RemoveStringSlice(target []string, start int, end int) []string {
	if end > cap(target)-1 {
		end = cap(target) - 1
	}
	if start > end {
		return make([]string, 0)
	}
	target = append(target[:start], target[end+1:]...)
	return target
}

func InsertStringSlice(target []string, index int, insert []string) []string {
	// 声明一个可容纳插入后全部数据的切片
	result := make([]string, len(target)+len(insert))
	// insert 需要插入到target index下标之后 分三段复制
	// 把target[:index]切片的数据复制到result
	// 复制insert到index之后
	// 复制target[index:]到index+len(insert)之后
	copy(result, target[:index])
	copy(result[index:], insert)
	copy(result[index+len(insert):], target[index:])
	return result
}

func Filter(s []int, fn func(int) bool) []int {
	var result []int

	for _, num := range s {
		if fn(num) {
			result = append(result, num)
		}
	}

	return result
}
