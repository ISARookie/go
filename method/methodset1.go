package main

import "fmt"

type List []int

func (l *List) Append(i int) {
	*l = append(*l, i)
}

func (l List) len() int {
	return len(l)
}

func main() {
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, len(lst))

	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.len())
}
