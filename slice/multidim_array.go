package main

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

// main /** 多维数组声明和初始化
func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
			//fmt.Printf("screen[%d][%d] value:%d\n", x, y, screen[x][y])
		}
	}
}
