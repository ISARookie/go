package main

var b = "G"

func main() {
	n()
	m()
	n()
}

func m() {
	b = "O"
	print(a)
}

func n() {
	print(a)
}
