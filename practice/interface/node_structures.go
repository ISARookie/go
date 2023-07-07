package main

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) setData(data interface{}) {
	n.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.setData("root node")
	a := NewNode(nil, nil)
	a.setData("left node")
	b := NewNode(nil, nil)
	b.setData("right node")
	root.le = a
	root.ri = b
	fmt.Printf("%v\n", root)
}
