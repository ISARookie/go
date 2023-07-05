package main

import (
	"example.com/m/v2/method/person"
	"fmt"
)

func main() {
	p := new(person.Person)
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName())
}
