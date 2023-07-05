package main

import "fmt"

type VCard struct {
	Name     string
	Birthday string
	Image    string
	Address  *Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

func main() {
	address := &Address{"street", "city", "country"}

	card := VCard{"name", "birthday", "image", address}

	(*card.Address).Country = "Street"

	fmt.Println(card.Address.Street)
}
