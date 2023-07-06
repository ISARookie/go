package main

import "fmt"

type List []int

func (l *List) Len() int {
	return l.Len()
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(appender Appender, start, end int) {
	for i := start; i <= end; i++ {
		appender.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(lener Lener) bool {
	return lener.Len()*10 > 42
}

func main() {
	var lst List
	fmt.Printf("lst %T\n", lst)
	CountInto(&lst, 1, 10)
	//if LongEnough(&lst) {
	//	fmt.Printf("- lst is long enough\n")
	//}
	plst := new(List)
	fmt.Printf("plst %T", plst)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}
