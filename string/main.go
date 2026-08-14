package main

import (
	"fmt"
)

//! 1
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


//! 2
func  ReverseString(s []string) []string {
	result := []string{}
	for i := len(s)-1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

//! 3

func  IsPalindrome(str string) bool {
	found := false
	left, right := 0, len(str)-1


	for left <= right {
		if str[left] == str[right] {
			left++
			right--
			found = true
		} else { 
			found = false
			return found
		}
	}
	return found
}
//! 4 

func strStr(haystack string, needle string) int {
	count := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[count] {
			count++
			if count == len(needle) {
				return i - count + 1
			}
		} else {
			count = 0
		}
	}
	return -1
}
func main() {
	 
	haystack := "aabaab"
	needle := "aab"
	resultPalindrome := strStr(haystack, needle)
	fmt.Println(resultPalindrome)


	// palindrome := "a"
	// resultPalindrome := IsPalindrome(palindrome)
	// fmt.Println(resultPalindrome)

	// word := []string{"H","a","n","n","a","h"}
	// result := ReverseString(word)
	// fmt.Println(result)



	// str := "((()))"
	// resultValid := Is_valid(str)
	// fmt.Println(resultValid)



}