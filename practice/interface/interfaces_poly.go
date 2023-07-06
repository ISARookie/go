package main

//import "fmt"
//
//type Shaper interface {
//	Area() float32
//}
//
//type Square struct {
//	side float32
//}
//
//func (sq *Square) Area() float32 {
//	return sq.side * sq.side
//}
//
//type Rectangle struct {
//	length, width float32
//}
//
//func (r Rectangle) Area() float32 {
//	return r.width * r.length
//}
//
//type Circle struct {
//	radius float32
//}
//
//func (c *Circle) Area() float32 {
//	return c.radius * c.radius
//}
//
//func main() {
//	r := Rectangle{5, 2}
//	q := &Square{5}
//	cl := &Circle{5}
//	shapes := []Shaper{r, q, cl}
//	fmt.Println("Looping through shapes for area ...")
//	for n, _ := range shapes {
//		fmt.Println("Shape details: ", shapes[n])
//		fmt.Println("Area of this shape is: ", shapes[n].Area())
//	}
//}
