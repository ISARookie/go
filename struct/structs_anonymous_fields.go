package main

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b   int
	c   float32
	in1 string
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 2
	outer.innerS.in1 = 4
	outer.in2 = 4
}
