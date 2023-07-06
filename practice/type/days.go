package main

import "fmt"

type Day int

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

var dayNames = [...]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func (d Day) String() string {
	if d >= MO && d <= SU {
		return dayNames[d]
	}
	return "Unknown"
}

func main() {
	fmt.Println(MO.String()) // Output: Monday
	fmt.Println(SU.String()) // Output: Sunday
}
