package main

type Point struct {
	x, y float64
}

type NamePoint struct {
	Point
	name string
}

//func (p *NamePoint) Abs() float64 {
//	return math.Sqrt(p.x*p.x + p.y*p.y)
//}

//func (n *NamePoint) Abs() float64 {
//	return n.Point.Abs() * 100
//}

func main() {
	//n := &NamePoint{Point{3, 4}, "Pythagoras"}
	//fmt.Println(n.Abs())
}
