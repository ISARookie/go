package main

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 0; i < 4; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
}
