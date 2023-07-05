package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Polar struct {
	Radius, Angle, Height float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(factor float64) {
	p.X = p.X * factor
	p.Y *= factor
}

func main() {
	point := Point{X: 3, Y: 4}
	fmt.Println("Point:", point)

	abs := point.Abs()
	fmt.Println("Abs:", abs)

	point.Scale(2)
	fmt.Println("Scaled Point:", point)
}
