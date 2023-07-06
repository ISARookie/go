package main

import "fmt"

type Log struct {
	msg string
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}

//func (c *Customer) Log() *Log {
//	return c.log
//}

//type Customer struct {
//	Name string
//	log  *Log
//}

type Customer struct {
	Name string
	Log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	//c.log = new(Log)
	//
	//c.log.msg = "1 - Yes we can"
	c = &Customer{"Barak Obama", Log{"1 - Yes we can"}}
	//c.Log().Add("2 - After me the world will be a better place!")

	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
	fmt.Println(&(c.Log))
}
