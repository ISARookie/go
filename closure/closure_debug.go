package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	where := func() {
		num, file, line, flag := runtime.Caller(1)
		log.Printf("%d:%s:%d:%t", num, file, line, flag)
	}
	where()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("spend time on where:%s\n", delta)
}
