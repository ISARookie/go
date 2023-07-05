package main

import "fmt"

type Engine struct {
	FuelType   string
	Horsepower int
}

func (e *Engine) Start() {
	fmt.Println("Engine started!")
}

type Car struct {
	wheelCount int
	engine     Engine
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

type Mercedes struct {
	Car
	int
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi, Merkel!")
}

func main() {
	engine := Engine{
		FuelType:   "Gasoline",
		Horsepower: 200,
	}

	car := Car{
		wheelCount: 4,
		engine:     engine,
	}

	mercedes := Mercedes{
		Car: car,
		int: 110,
	}

	fmt.Println("Number of Wheels:", mercedes.numberOfWheels())

	mercedes.engine.Start()

	mercedes.sayHiToMerkel()
}
