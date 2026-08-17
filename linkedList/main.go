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

func (l *Head) delete(value int)  { 
	if l.Head != nil {
		if l.Head.value == value {
			l.Head = l.Head.next
			return
		}
	}
		
	current := l.Head
	if current.next != nil {
		for current.next.value != value {
			current = current.next
		}
		current.next = current.next.next
	}
	
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
	node.insert(10)
	node.insert(20)
	node.insert(30)
	node.insert(40)
	node.insert(50)
	node.insert(60)
	node.delete(50)
	node.print()

}
