package main

import "fmt"


type Node struct { 
	Val int
	Next *Node
}

type Head struct { 
	Head *Node
}


func (h *Head) insert(Val int) {
	ValNode := &Node{Val: Val, Next: nil}

	if h.Head == nil {
		h.Head = ValNode
		return
	}

	current := h.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = ValNode
}

func (l *Head) delete(Val int)  { 
	if l.Head != nil {
		if l.Head.Val == Val {
			l.Head = l.Head.Next
			return
		}
	}
		
	current := l.Head
	if current.Next != nil {
		for current.Next.Val != Val {
			current = current.Next
		}
		current.Next = current.Next.Next
	}
	
}

func (l *Head) print() {
	current := l.Head
	for current != nil { 
		fmt.Println(current)
		current = current.Next
	}
}




//! 1

func MiddleNode(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}


//! 2


func RemoveDuplicatesFromLinkedList(head *Node) *Node {

	k := head.Val
		for head != nil && head.Next != nil{
			if k == head.Next.Val {
				head.Next = head.Next.Next
			}else {
				head = head.Next
				k = head.Val
			}
		}

	return head
}
//! 3- -> work

func removeElements(head *Node, val int) *Node {
    for head != nil && head.Val == val {
        head = head.Next
    }

    curr := head
    for curr != nil && curr.Next != nil {
        if curr.Next.Val == val {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    return head
}

//! 4


func Reverse(head *Node) *Node {
	var prev *Node = nil
	curr := head
	for curr != nil { 
 		next := curr.Next   
        curr.Next = prev    	
        prev = curr         
        curr = next 
	}
	return prev
}

//! 5 -> not full

func deleteDuplicates(head *Node) *Node {
	result := head
    k := head.Val
	for head != nil && head.Next != nil {
		if k == head.Next.Val { 
				head.Next = head.Next.Next
		} else { 
			head = head.Next
			k = head.Val
		}
	}
	return result
}

//! 6 -> it didnt work 

func deleteNode(node *Node) {
	fmt.Println(node)
	if node != nil {
		if node.Val == 5 {
			node = node.Next
			return
		}
	}
	current := node
	for current != nil && current.Next != nil{
		if current.Next.Val == 5 {
			current.Next = current.Next.Next
			return
		} else {
			 current = current.Next
		}
	}

}

//! 7

func deleteMiddle(head *Node) *Node {
	if head.Next == nil  {
		return nil
	}
	var prev *Node = nil
    slow, fast := head,head 
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next

	}
	prev.Next = slow.Next
	return head
}

//! 8 it didint work 

func removeNthFromEnd(head *Node, n int) *Node {
	    
	count := 0 
	prev := head

	if head.Next == nil || head == nil {
		return nil
	}
	
	for head != nil { 
		if count > n { 
			prev = prev.Next
		}
		head = head.Next
		count++
	}
	prev.Next = prev.Next.Next
	return head
}
//! 9 -> beats 92%

func modifiedList(nums []int, head *Node) *Node {
    m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {	
		m[nums[i]] = true
	}
	for head != nil && m[head.Val] == true {
		head = head.Next
	}
	curr := head
	for curr != nil && curr.Next != nil { 
		if m[curr.Next.Val] == true {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func main() {
	head := Head{}
	
	head.insert(10)
	head.insert(2)
	head.insert(3)
	head.insert(4)
	head.insert(5)
	head.insert(6)
	head.insert(7)

	nums := []int{1,2,3,4}
	head.Head = modifiedList(nums, head.Head)



	// result := Reverse(head.Head)
	// head.Head = deleteDuplicates(head.Head)

	// deleteNode(head.Head)
	// deleteMiddle(head.Head)

	// val := 7
	// head.Head = removeElements(head.Head, val)


	head.print()


	// n := 2
	// removeNthFromEnd(head.Head, n)




	// MiddleNode(head.Head)
	// RemoveDuplicatesFromLinkedList(head.Head)

	







}
