package main

import "fmt"

type Rectangle struct {
	WIDTH, HEIGHT int
}

func Area(r *Rectangle) int {
	return r.HEIGHT * r.WIDTH
}

func Primeters(r *Rectangle) int {
	return r.HEIGHT*2 + r.WIDTH*2
}
func main() {
	rectangle := Rectangle{10, 11}
	fmt.Println(Area(&rectangle))
	fmt.Println(Primeters(&rectangle))
}
