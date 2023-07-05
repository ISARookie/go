package main

import "fmt"

type employee struct {
	salary float64
}

func (e *employee) giveRaise(f float64) *employee {
	e.salary += e.salary * f
	return e
}

func main() {
	emp := employee{100.}
	fmt.Printf("Add raise of salary: %v\n", *emp.giveRaise(0.2))
}
