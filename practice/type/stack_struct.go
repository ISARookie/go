package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack struct {
	ix   int
	data [LIMIT]int
}

func main() {
	st1 := new(Stack)
	fmt.Printf("%v\n", st1)
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	st1.Push(99)
	fmt.Printf("%v\n", st1)
	st1.Push(999)
	fmt.Printf("%v\n", st1)
}

func (st *Stack) Push(n int) {
	if st.ix+1 > LIMIT {
		return
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	st.ix--
	return st.data[st.ix]
}

func (st Stack) String() string {
	str := ""
	for i := 0; i < st.ix; i++ {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(st.data[i]) + "]"
	}
	return str
}
