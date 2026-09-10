package main 

import ("fmt")


// Node
type Node struct { 
	key int
	left *Node
	right *Node 
}
// Insert
func (n *Node) insert(val int) { 
	if val > n.key {
		if n.right == nil {
			n.right = &Node{key: val}
		} else { 
			n.right.insert(val)
		}
	} else if val < n.key {
		if n.left == nil {
			n.left = &Node{key: val}
		} else { 
			n.left.insert(val)
		}
	}
}

func (n *Node) search(val int) bool { 
	if n == nil {
		return false
	}
	if val > n.key {
		return n.right.search(val)
	} else if val < n.key { 
		return n.left.search(val)
	}

	return true
}
// Search



func main() { 
	tree := &Node{key: 100}
	tree.insert(101)
	tree.insert(104)
	fmt.Println(tree)
	fmt.Println(tree.search(100))

}