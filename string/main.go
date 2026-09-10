package main

import (
	"fmt"
	"strings"
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

//! 10

func nextGreatestLetter(letters []byte, target byte) byte {
    
	for i := 0; i < len(letters); i++ {
		if target < letters[i] {
			return letters[i]
		}
	}
	return letters[0]
}

//! 11

func longestCommonPrefix(strs []string) string {
	count := 0

	if len(strs) == 0 {
		return ""
	}

	result := len(strs[0])

	for i := 0; i < len(strs); i++ {
		result = min(len(strs[i]), result)
	}

	for i := 0; i < result; i++ {
		max := 0

		for k := 0; k < len(strs)-1; k++ {
			if strs[k][i] == strs[k+1][i] {
				max++

				if max == len(strs)-1 {
					count++
				}
			} else {
				return strs[0][:count]
			}
		}
	}

	return strs[0][:count]
}

//! 12

func isIsomorphic(s string, t string) bool {
    m1 := make(map[byte]byte)
    m2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if _, ok := m1[s[i]]; !ok {
			m1[s[i]] = t[i]
		}
		if _, ok := m2[t[i]]; !ok {
			m2[t[i]] = s[i]
		}
		
		if !(m1[s[i]] == t[i] && m2[t[i]] == s[i]) {
			return false
		}
	}

	return true
}

//! 13 too easy

func toLowerCase(s string) string {
	toLower := ""
	for i := 0; i < len(s); i++ {
		toLower += strings.ToLower(string(s[i]))
	}
	return toLower
}

//! 14 too easy -> time O(n), space O(1)

func mostWordsFound(sentences []string) int {
	maximum := 0

    for i := 0; i < len(sentences); i++ {
		count := 0
		for k := 0; k < len(sentences[i]); k++ {
			if sentences[i][k] == ' ' {
				count++
			}
		}
		maximum = max(maximum, count)
	}
	return maximum + 1
}

//! 15 

func findWordsContaining(words []string, x byte) []int {
    containsX := []int{}
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			if words[i][k] == x {
				containsX = append(containsX, i)
				break
			}
		}
	}
	return containsX
}
//! 16
func countKeyChanges(s string) int {
	count := 0
    for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			if !(s[i] - s[i+1] == 32 || s[i+1] - s[i]  == 32) {
				count++
			}
			
		}
	}
	return count
}

//! 17



func main() {



	// s := "AabBcC"
	// result := countKeyChanges(s)
	// fmt.Println(result)



	// words := []string{"leet","code"}
	// x := byte('e')
	// result := findWordsContaining(words, x)
	// fmt.Println(result)

	// sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	// result := mostWordsFound(sentences)
	// fmt.Println(result)
		


	// s := "LOVELY"
	// result := toLowerCase(s)
	// fmt.Println(result)
		

	// s := "bbbaaabe"
	// t := "aaabbbba"
	// result := isIsomorphic(s, t)
	// fmt.Println(result)
	



	// words := []string{"a"}
	// result := longestCommonPrefix(words)
	// fmt.Println(result)
	
	

	// s := []byte{'d','f','j'}
	// target := 'a'
	// result := nextGreatestLetter(s, target)
	// fmt.Println(result)


	// s := "   fly me   to   the moon  "
	// result := lengthOfLastWord(s)
	// fmt.Println(result)

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