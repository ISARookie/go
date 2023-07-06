package main

import "fmt"

type Base struct{}

// Magic /**
func (Base) Magic() {
	fmt.Println("base magic")
}

func (receiver Base) MoreMagic() {
	receiver.Magic()
	receiver.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
