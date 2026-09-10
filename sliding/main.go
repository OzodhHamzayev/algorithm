package main

import ("fmt")

//! 1
func maxSum(nums []int, k int) int { 
	sum := 0
	maxSum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if i >= k-1 {
			if maxSum < sum {
				maxSum = sum
			}

			sum -= nums[i-k+1]
		}
	}

	return maxSum
}
//! 2
func maxSumMultiply(nums []int, k int) int { 
	sum := 1
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		sum *= nums[i]
		if i >= k-1 {
			if maxSum < sum {
				maxSum = sum
			}
			sum /= nums[i-k+1]
		}
	}
	return maxSum
}
//! 3
func fixedSize(nums []int, k int, target int) int { 
	sum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i >= k-1 {
			if sum == target {
				count++
			}
			sum -= nums[i-k+1]
		}
	}
	return count
}
//! 4

func substringAnagram(word1 string, word2 string) bool { 
	m := make(map[byte]int)
	count := 0
	for i := 0; i < len(word1); i++ {
		m[word1[i]]++
	}
	for i := 0; i < len(word2); i++ {
		if m[word2[i]] >= 1 {
			m[word2[i]]--
			count++
		}
		if count == len(word2) {
			return true
		}
	}
	return false
}

//! 5

func addNums(nums []int, target int) bool {
	shrink := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum > target { 
			sum -= nums[shrink]
			shrink++
		}
		if sum == target {
			return true
		}
	}
	return false
}
//! 6
func longestSubarray(nums []int, target int) int {
	shrink := 0
	sum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum > target { 
			sum -= nums[shrink]
			shrink++
		}
		if sum == target {
			if i-shrink+1 > count {
				count = i-shrink+1
			}
		}
	}
	return count
}

//! 7
func isVowels(c byte) bool { 
	return c == 'a' ||  c == 'e' ||c == 'i'||c == 'u'||c == 'o'
}

func maxVowels(s string, k int) int {
	left, right := 0, 0
    maximum := 0
	count := 0
	for right < len(s) {
		if right-left+1 <= k {
			if isVowels(s[right]) {
				count++
			}
			right++
			maximum = max(maximum, count)
		} else { 
			if isVowels(s[left]) {
				count--	
			}
			left++
		}
	}
	return maximum
}


//! 8

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	minimum := 0
	shrink := 0
    for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			if minimum == 0 {
				minimum = i-shrink+1
			}else { 
				minimum = min(i-shrink+1, minimum)
			}
			sum -= nums[shrink]
			shrink++
		}
	}
	return minimum
}


//! 9

func longestOnes(nums []int, k int) int {
    count := 0
	shrink := 0
	maximum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		}
		for count > k {
			if nums[shrink] == 0 {
				shrink++
				count--
			}else {
				shrink++
			}
		}	
		maximum = max(i-shrink+1, maximum)
	}
	return maximum
}

//! 10

func maximumSubarraySum(nums []int, k int) int64 {
    shrink := 0
	maximum, sum := 0, 0
	
	m := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		for m[nums[i]] {
			sum -= nums[shrink]
			delete(m, nums[shrink])
			shrink++
		}
		m[nums[i]] = true
		sum += nums[i]
		if i-shrink+1 > k {
			sum -= nums[shrink]
			delete(m, nums[shrink])
			shrink++
		}

		if i-shrink+1 == k {
			maximum = max(sum, maximum)
		}
		
	}
	return int64(maximum)
}
//! 11


func main() { 


	// nums := []int{1,5,4,2,9,9,9,1,4}
	// k := 3
	// result := maximumSubarraySum(nums,k)
	// fmt.Println(result)


	// nums := []int{1,1,1,0,0,0,1,1,1,0,1}
	// k := 2
	// result := longestOnes(nums,k)
	// fmt.Println(result)


	// k := 7
	// nums := []int{2,3,1,2,4,3}
	// result := minSubArrayLen(k, nums)
	// fmt.Println(result)

	// s := "abciiidef"
	// k := 3
	// result := maxVowels(s, k)
	// fmt.Println(result)


	// nums := []int{17,1,1,1,1,1,11,1,0,0,0,0,0,0,0,0,0,0,0}
	// target := 17
	// result := longestSubarray(nums, target)
	// fmt.Println(result)

	// nums := []int{2,3,2,19,3,10,3,8,5,0,4,4}
	// target := 17
	// result := addNums(nums, target)
	// fmt.Println(result)

	// word1 := "greyhundso"
	// word2 := "hoy"
	// result := substringAnagram(word1, word2)
	// fmt.Println(result)

	// nums := []int{2,3,2,2,3,1,3,8,5,0,2,4}
	// k := 3
	// target := 7
	// result := fixedSize(nums, k, target)
	// fmt.Println(result)



	// nums := []int{1,4,1,6,-3,3,-5,2,26}
	// k := 4
	// result := maxSumMultiply(nums, k)
	// fmt.Println(result)


	// nums := []int{1,4,1,10,25,3,5,0,26}
	// k := 4
	// result := maxSum(nums, k)
	// fmt.Println(result)
}
