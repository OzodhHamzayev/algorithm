package main

import ("fmt")


//! 1

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	left := 0
	count := 0
	for i := 0; i < len(s); i++ {
		for m[s[i]] {
			delete(m, s[left])
			left++
		}
			m[s[i]] = true
			count = max(count,i-left+1)	 
	}
	return count
}

//! 2

func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if m[ransomNote[i]] > 0 {
			m[ransomNote[i]]--
		} else {
			return false
		}
	}
	return true
}

//! 3

func IsAnagram(s string, t string) bool {
    m := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

    for i := 0; i < len(s); i++ { 
      m[s[i]]++
    }
    for i := 0; i < len(t); i++ { 
      if m[t[i]] > 0 { 
        m[t[i]]--
      } else { 
        return false
      }
    }
    return true
}

//! 4

func FirstUniqChar(s string) int {
    m := make(map[byte]int) 
    for i := 0; i < len(s); i++ { 
      m[s[i]]++
    }
    for i := 0; i < len(s); i++ { 
      if m[s[i]] == 1 { 
        return i
      }
    }
  return -1
}

//! 5

func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
		m[nums[i]] = true
	}
	for i := 0; i < len(nums); i++ {
		if m[min] {
			count += 1
			min += 1
		}else { 
			return count
		}
	}
	return count
}

//! 6

func uniqueOccurrences(arr []int) bool {
    
}

func main() {

	nums := []int{0,0,1,2,3,4,5,6,7,8}
	result := longestConsecutive(nums)
	fmt.Println(result)




	// s := "leetcode"
	// result := FirstUniqChar(s)
	// fmt.Println(result)



	// s := "aabbcc"
	// t := "aabbc"
	// result := IsAnagram(s, t)
	// fmt.Println(result)



	// ransomNote := "aabbcc"
	// magazine := "aabbc"
	// result := canConstruct(ransomNote, magazine)
	// fmt.Println(result)




	// word := "abbcebb"
	// result := lengthOfLongestSubstring(word)
	// fmt.Println(result)
}
