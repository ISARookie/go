package main

//
//import "fmt"
//
//type Base struct {
//	id string
//}
//
//type Person struct {
//	Base
//	FirstName string
//	LastName  string
//}
//
//type Employee struct {
//	Person
//	salary int
//}
//
//func (b *Base) SetId(s string) {
//	b.id = s
//}
//
//func (e Base) Id() string {
//	return e.id
//}
//
//func main() {
//	e := new(Employee)
//	e.id = "id"
//	e.FirstName = "FN"
//	e.LastName = "LN"
//	e.salary = 100
//	e.Id()
//	e.SetId("ID")
//	fmt.Println(e)
//}
