package main

import (
	"strconv"
)

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + "°C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func print(args ...interface{}) {
	//for i, arg := range args {
	//	if i > 0 {
	//		os.Stdout
	//		_, err := os.Write([]byte("hello golang"))
	//	}
	//	switch a := arg.(type) {
	//	case Stringer:
	//		os.Stdout.
	//	}
	//
	//}
}
