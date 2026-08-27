package main

import "fmt"


type Node struct { 
	Val int
	Next *Node
}

type Head struct { 
	Head *Node
}


func (h *Head) insert(val int) {
	ValNode := &Node{Val: val, Next: nil}

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

//! 8 it work 

func removeNthFromEnd(head *Node, n int) *Node {
	dummy := &Node{Next: head}
	left, right := dummy, head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next

	return dummy.Next
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



//! 11


func IsPalindrome(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil { 
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow = ReverseSlow(slow)

	for slow != nil {
		if slow.Val == head.Val {
			slow = slow.Next
			head = head.Next
		}else {
			return false
		}
	}
	return true
}
func ReverseSlow(head *Node) *Node { 
	var prev *Node = nil
	curr := head
	for curr != nil { 
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}


//! 12

func MergeTwoLists(head1 *Node, head2 *Node) *Node {
	dummy := &Node{}
	current := dummy
	for head1 != nil && head2 != nil { 
		if head1.Val < head2.Val {
			current.Next = head1
			current = current.Next
			head1 = head1.Next
		} else {
			current.Next = head2
			current = current.Next
			head2 = head2.Next
		}
	}
	if head1 != nil {
		current.Next = head1
	} else {
		current.Next = head2
	}
	return dummy.Next

	
}
//!------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
//! 13
func hasCycle(head *Node) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            return true
        }
    }
    return false
}
//! 14



//! 28
//!------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
func main() {
	head := Head{}


	head.insert(1)
	head.insert(2)
	

	
	// result := hasCycle(head.Head)
	// fmt.Println(result)


	// a := IsPalindrome(head.Head)
	// fmt.Println(a)



	// nums := []int{1,2,3,4}
	// head.Head = modifiedList(nums, head.Head)



	// result := Reverse(head.Head)

	// deleteNode(head.Head)
	// deleteMiddle(head.Head)

	// val := 7
	// head.Head = removeElements(head.Head, val)




	// n := 2
	// head.Head = removeNthFromEnd(head.Head, n)
	// head.print()
	// head2.print()





	// MiddleNode(head.Head)
	// RemoveDuplicatesFromLinkedList(head.Head)

	







}
