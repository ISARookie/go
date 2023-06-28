package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)

	for element := l.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

}
