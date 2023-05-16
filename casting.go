package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 10, 0

	b += a
	c := a / b
	print(c)

	//type BigFlag int
	//const (
	//	Active BigFlag = 1 << iota
	//	Send
	//	Receive
	//)
	//
	//flag := Active | Send
	//fmt.Println(flag)

	//type ByteSize float64
	//const (
	//	_           = iota // 通过赋值给空白标识符来忽略值
	//	KB ByteSize = 1 << (10 * iota)
	//	MB
	//	GB
	//	TB
	//	PB
	//	EB
	//	ZB
	//	YB
	//)
	//fmt.Println(KB)
	//fmt.Println(MB)

	/*
		complex64 32位实数和虚数
		complex128 64位实数和虚数
	*/

	//var Pi = 3.1415926
	//var n int16 = 34
	//var m int32
	//
	//m = int32(n)
	//
	//fmt.Printf("32 bit int is：%d \n", m)
	//fmt.Printf("64 bit int is：%d \n", n)
	//fmt.Printf("%4.2f", Pi)
}

func Uint8FormInt(n int) (uint8, error) {
	// 验证取值范围
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	//panic(fmt.Sprintf("%g is out of the int32 range", x))
	return 0
}
