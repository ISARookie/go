package main

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func main() {
	celsius := Celsius(22.2222)
	s := celsius.String()
	fmt.Println(s)
}
