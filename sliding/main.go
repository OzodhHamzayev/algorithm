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
func main() { 


	nums := []int{17,1,1,1,1,1,11,1,0,0,0,0,0,0,0,0,0,0,0}
	target := 17
	result := longestSubarray(nums, target)
	fmt.Println(result)

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
