package main

import (
	"fmt"
)


func main() {
	 
	word := []string{"H","a","n","n","a","h"}
	result := ReverseString(word)
	fmt.Println(result)

	str := "((()))"
	resultValid := Is_valid(str)
	fmt.Println(resultValid)



}

//! 3


//! 4
func Is_valid(s string) bool {
	count1:= 0
	count2 := 0
	leng := len(s)-1
	if s[leng] == 40 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			count1++
		} else if s[i] == 41 {
			count2++
		}
	}
	if count1 == count2 {
		return true
	}
	return false
}


//! 5
func  ReverseString(s []string) []string {
	result := []string{}
	for i := len(s)-1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}
