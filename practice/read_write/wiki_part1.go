package main

import (
	"bufio"
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() {
	outputFile, _ := os.OpenFile(p.Title, os.O_CREATE|os.O_WRONLY, 0)
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)
	writer.Write(p.Body)
	writer.Flush()
	//writer.Write(p.Body)
	//writer.WriteString("p.Body")
	//writer.Flush()
	//outputFile.Write(p.Body)
}

func (p *Page) load() {
	file, err := os.ReadFile(p.Title)
	if err != nil {
		panic("FILE NOT FOUND")
	}
	fmt.Println(string(file))
}

func main() {
	p := new(Page)
	p.Title = "test2.text"
	p.Body = []byte("hello world!")
	p.save()
	p.load()
}
