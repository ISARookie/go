package main

//
//import (
//	"fmt"
//	"strconv"
//)
//
//const LIMIT = 4
//
//type Stack [LIMIT]int
//
//func (st *Stack) Push(n int) {
//	for ix, v := range st {
//		if v == 0 {
//			st[ix] = n
//			break
//		}
//	}
//}
//
//func (st *Stack) Pop() int {
//	v := 0
//	for ix := len(st) - 1; ix >= 0; ix-- {
//		if v = st[ix]; v != 0 {
//			st[ix] = 0
//			return v
//		}
//	}
//	return 0
//}
//
//func (st Stack) String() string {
//	str := ""
//	for ix, v := range st {
//		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "]"
//	}
//	return str
//}
//
//// Stack /** 先入后出，设置LIMIT控制容量
//func main() {
//	st1 := new(Stack)
//	fmt.Printf("%v\n", st1)
//	st1.Push(3)
//	fmt.Printf("%v\n", st1)
//	st1.Push(7)
//	fmt.Printf("%v\n", st1)
//	st1.Push(10)
//	fmt.Printf("%v\n", st1)
//	st1.Push(99)
//	fmt.Printf("%v\n", st1)
//	p := st1.Pop()
//	fmt.Printf("Popped %d\n", p)
//	fmt.Printf("%v\n", st1)
//}
