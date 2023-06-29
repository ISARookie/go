package main

import (
	"fmt"
	"time"
)

func main() {
	ISAM()
}

func ISAM() {
	now := time.Now()
	hour := now.Hour()
	if hour > 12 {
		IsEvening()
	} else {
		IsAfternoon()
	}
}

func IsAfternoon() {
	fmt.Println("Good Day")
}

func IsEvening() {
	fmt.Println("Good Night")
}
