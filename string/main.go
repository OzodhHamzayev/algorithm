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


//! 5

func firstUniqChar(s string) int {
	m := make(map[byte]int, 26)

    for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i	
		}	 
	}
	return 	-1
}

//! 6

func numJewelsInStones(jewels string, stones string) int {
    m := make(map[byte]bool)
	count := 0
	for i := 0; i < len(jewels); i++ {
		m[jewels[i]] = true
	}
	for i := 0; i < len(stones); i++ {
		if m[stones[i]] == true {
			count++
		}
	}
	return count
}

//! 7

func isVowel(c byte) bool {
	return  c == 'a' || c == 'e' || c == 'i' ||
        c == 'o' || c == 'u' ||
        c == 'A' || c == 'E' || c == 'I' ||
        c == 'O' || c == 'U'
}

func reverseVowels(s string) string {
	chars := []byte(s)
    left, right := 0, len(s)-1
	for left < right { 
		if !isVowel(s[left]) {
			left++
		}else if !isVowel(s[right]) { 
			right--
		} else if isVowel(s[left]) && isVowel(s[right]) {
			chars[left], chars[right] = chars[right], chars[left]	
			left++
			right--
		}
	}
	return string(chars)
}
//! 8

func mergeAlternately(word1 string, word2 string) string {
    i := 0
	result := ""
	for i < len(word1) && i < len(word2) {
		result += string(word1[i])
		result += string(word2[i])
		i++
	}
	if i < len(word1) {
		result += string(word1[i:])
	} else if i < len(word2) { 
		result += string(word2[i:])
	}
	return result
}

//! 9


func lengthOfLastWord(s string) int {
	count := 0
	result := 0
    for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			count = 0
		}else { 
			count++
			result = count
		}

	}
	return result
}


func main() {

	s := "   fly me   to   the moon  "
	result := lengthOfLastWord(s)
	fmt.Println(result)

	// word1 := "ab"
	// word2 := "pqrs"
	// result := mergeAlternately(word1, word2)
	// fmt.Println(result)

	// word := "IceCreAm"
	// resultPalindrome := reverseVowels(word)
	// fmt.Println(resultPalindrome)

	// jewels := "aA"
	// stones := "aAAbbbb"
	// resultPalindrome := numJewelsInStones(jewels, stones)
	// fmt.Println(resultPalindrome)



	// s := "loveleetcode"
	// resultPalindrome := firstUniqChar(s)
	// fmt.Println(resultPalindrome)


	 
	// haystack := "aabaab"
	// needle := "aab"
	// resultPalindrome := strStr(haystack, needle)
	// fmt.Println(resultPalindrome)


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