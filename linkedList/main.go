package main

import "fmt"


type Node struct { 
	value int
	next *Node
}

type Head struct { 
	Head *Node
}


func (h *Head) insert(value int) {
	valueNode := &Node{value: value, next: nil}

	if h.Head == nil {
		h.Head = valueNode
		return
	}

	current := h.Head
	for current.next != nil {
		current = current.next
	}
	current.next = valueNode
}

func (h *Head) print() {
	current := h.Head
	for current != nil { 
		fmt.Println(current)
		current = current.next
	}
}

func main() {
	node := Head{}
	node.insert(12)
	node.insert(11)
	node.insert(10)
	node.insert(9)
	node.insert(9)
	node.insert(7)
	node.print()

}
