package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastIndex := len(s.items) - 1
	popValue := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return popValue, true
}

func (s *Stack) top() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

//! 1

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if top != m[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}


//! 2 it didint work (hard)

func longestValidParentheses(s string) int {
    m := map[byte]byte{
		')': '(',
	}
	stack := []byte{}
	count := 0
	for i := 0; i < len(s); i++ {
		bracket := s[i]
		if bracket == '(' {
			stack = append(stack, bracket)
		}else { 
			if len(stack) == 0 {
				continue
			}
			top := stack[len(stack)-1]
			if top == m[bracket] {
				count += 2
				stack = stack[:len(stack)-1]
			}

		}
	}
	return count
}


//! 3

func removeStars(s string) string {
     stack := []byte{}

	for i := 0; i < len(s); i++ {
		word := s[i]
		if  word != '*' {
			stack = append(stack, word)
		} else { 

			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func main() {


	s := "aaa*b"
	result := removeStars(s)
	fmt.Println(result)



	// s := "()(()"
	// result := longestValidParentheses(s)
	// fmt.Println(result)

	// s := "]"
	// result := isValid(s)
	// fmt.Println(result)

	// stack := Stack{}

	// stack.push(10)
	// stack.push(20)
	// stack.push(20)
	// stack.pop()
	// stack.top()

	// fmt.Println(stack)

}
