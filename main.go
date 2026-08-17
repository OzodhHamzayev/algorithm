package main

import (
	"fmt"
	"unicode"
)

//! 1
// TODO: 2 ta loop ishlatganimning sababi — value[0] va value[10]  dagi sonlarning yig‘indisi target ga teng bo‘lishi mumkin.
// TODO Ular ketma-ket kelmasligi mumkin. Agar kerakli sonlar ketma-ket  kelganida, bitta loopdan foydalanardim.
func TwoNumberSum(array []int, target int) []int {
	lens := len(array)-1
	if lens == 0 {
		return []int{}
	}
	for i := 0; i < len(array); i++ {
		for k := i+1; k < len(array); k++ {
			if array[i] + array[k] == target {
				if array[i] > array[k] {
					return []int{array[k], array[i]}
				}
					return []int{array[i], array[k]}
			} 
		}
	}
	return []int{}
}


//! 2
//TODO agar bizning index == 0 ga boladigan bolsa indexni joyini keyingi keladigan index bilan change qiladi va bu xolat bir 
//TODO necha marta davom etgani uchun 2ta loop qildim 
func MoveZeroes(nums []int) []int {
  for i := 0; i < len(nums)-1; i++ {
    for k := 0; k < len(nums)-1; k++ {
      if nums[k] == 0 {
        nums[k], nums[k+1] = nums[k+1], nums[k]	
      }
    }
	}
	return nums
}


//! 3
// TODO: arrayda taxminan 20 ta son bor. sequence arrayidagi qiymatlar  array ichida shu tartibda uchrashi kerak. 
// TODO: sequenceIndex = 0 dan boshlaymiz. Loop orqali arrayni tekshiramiz.  Agar arraydagi biror qiymat sequence[0] ga teng bo‘lsa, sequenceIndex++ qilamiz. 
// TODO: Keyingi iterationda sequence[1] bilan tekshiramiz. // Agar u ham mos kelsa, sequenceIndex yana oshadi va shu tarzda davom etadi.
func  IsValidSubsequence(array []int, sequence []int) bool {
	sequenceIndex := 0
	for i := 0; i < len(array); i++ {
		if sequenceIndex == len(sequence) {
			break
		}
		if array[i] == sequence[sequenceIndex] {
			sequenceIndex++
		}
	}
	return sequenceIndex == len(sequence)
	
}


//! 4
//TODO k -> nechta soni chap tomonga olib o'tishimizdagi qiymat. Biz birinchi loop da i < k qildik yani u faqat k marta loop
//TODO aylanishi kerak masalan bizda nums := []int{1,2,3,4,5,6,7} va k = 3. ichki loop orqaga 3 marta yuradi va kerakli sonlar
//TODO o'zgaradi
func Rotate(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		for j := len(nums)-1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
			
		}
	}
	return nums
}



//! 5
//TODO 2-masala bilan bir xil
func MoveElementToEnd(array []int, toMove int) []int {
	for i := 0; i < len(array)-1; i++ {
		for k := 0; k < len(array)-1; k++ {
			if array[k] == toMove {
				array[k], array[k+1] = array[k+1], array[k]
			}
		}
	}
	return array
}






//! 6 

func SortedSquares(nums []int) []int {
	numsSort := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			fmt.Println(nums[i])
			nums[i] = -nums[i]
			numsSort = append(numsSort, nums[i] * nums[i])
		} else {
			numsSort = append(numsSort, nums[i] * nums[i])
		}
		for k := 0; k < len(numsSort)-1; k++ {
			if numsSort[k] > numsSort[k+1] {
				numsSort[k], numsSort[k+1] = numsSort[k+1], numsSort[k]
			}
		}
	}
	return numsSort
}


//! 7


// func LeftRightDifference(nums []int) []int {
// 	result := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		left := 0
// 		right := 0

// 		for k := 0; k < len(nums); k++ {
// 			if k < i {
// 				left += nums[k]
// 			} else if k > i {
// 				right += nums[k]
// 			}
// 		}
// 		if left > right {
// 			result = append(result, -1)
// 		} else if left < right {
// 			result = append(result, 1)
// 		} else {
// 			result = append(result, 0)
// 		}
// 	}
// 	return result
// }

//! 8


func IsMonotonic(array []int) bool {

	if len(array)-1 == 0 {
		return true	
	}
	count := 0
	for i := 0; i < len(array)-1; i++ {
		if array[count] == array[count+1] {
			count++
		}else if array[count] < array[count+1] {
			if array[i] > array[i+1] {
				return false
			}
		} else if array[count] > array[count+1] { 
			if array[i] < array[i+1] {
				return false
			}
		}
	}
	return true
}


//! 9

func  BestSeat(seats []int) int {
	count := 0
	max := 0
	index := 0
	for i := 0; i < len(seats); i++ {
		if seats[i] != 1 {
			count++
		} else if seats[i] == 1 {
			if max < count {
				max = count
				index = i
			} 
				count = 0 
		}
	}

	if max == 0 {
		return -1
	}

	max, index = max/2, index-1
	result := index-max
	if seats[result] == 0 {
		seats[result] = 1
		return result
	}
	return -1
}

//! 10 ->


func removeDuplicates(nums []int) int {

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[k] == nums[i] {
			continue
		} else {
			nums[k+1], nums[i] = nums[i], nums[k+1]
			k = k+1
		}
	}
	return k+1
}




//! 11

func MoveZeroes2(nums []int) []int {
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
	return nums
}

//! 12

func MoveElementToEnd2(array []int, toMove int) []int {
	target := 0
	for i := 0; i < len(array); i++ {
		if array[i] != toMove {
			array[i], array[target] = array[target], array[i]
			target++
		}
	}
	return array
}

//! 12-2

func MoveElementToEnd3(array []int, toMove int) []int {
	left, mid, right := 0,0, len(array)-1
	for mid <= right {
		if array[mid] == toMove {
			mid++
		} else {
			array[mid], array[left] = array[left], array[mid]
			left++
			mid++
		}
	}
	return array
}


//! 13 

func MaxArea(nums []int) int {
	left, right := 0, len(nums)-1
	waterResult := 0
	for left <= right { 
		water := 0
		if nums[left] >= nums[right] {
			water = nums[right] * (right-left)
		} else {
			water = nums[left] * (right-left)
		}
		if waterResult < water {
		waterResult = water
		}
		if nums[left] > nums[right] {
			right--
		} else if nums[left] < nums[right] {
			left++
		} else {
			left++
			right--
		}
	}
	return waterResult
}

//! 17



func sortingElement(nums []int) []int {
	 low, high := 0, len(nums)-1
	 mid := (high+low)/2
	 for low <= high {
		if nums[mid] < nums[low] {
			nums[mid], nums[low] = nums[low], nums[mid]
			mid--
			fmt.Println(nums)

		} else if nums[mid] > nums[high] {
			nums[mid], nums[high] = nums[high], nums[mid]
			mid++
		} else {
			break
		}
	 }
	 return nums
}


//! 14 

  func FirstDuplicateValue(array []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(array); i++ {

		if m[array[i]] == true {
			return array[i]
		}
		m[array[i]] = true
		
	}
    return -1
  }

//! 15  -> o(n kvadrat)

func ZeroSumSubarray(nums []int) bool {
	for i := 0; i < len(nums); i++ {
	sum := 0
		for k := i; k < len(nums); k++ {
			sum += nums[k]
			if sum == 0 {
				return true
			}
		}
	}
	return false
}


//! 16 -> o(n kvadrat)


func  FirstNonRepeatingCharacter(str string) int {
	m := make(map[byte]bool)
	for i := 0; i < len(str); i++ {

		for k := i+1; k < len(str); k++ {
			if str[i] == str[k] {
				m[str[i]] = true
			} 
		}
		if m[str[i]] == false  {
			return i
		}
	}

	return -1
}
//! 17 -> o(n)
func  FirstNonRepeatingCharacter2(str string) int {
    m := make(map[byte]int)

    for i := 0; i < len(str); i++ {
        m[str[i]]++
    }

		for i := 0; i < len(str); i++ {
			if m[str[i]] == 1 {
				return i
			}
	}
	return -1
}

//! 18 

func LargestRange(array []int) []int {
	min := array[0]
	max := 0
	result := make([]int, 2)
	for i := 0; i < len(array); i++ {
		for k := 0; k < len(array); k++ {
			if min >= array[k] && i == 0 {
				min = array[k]
				result[0] = min
				max = min
				max++
				} else if max == array[k] {
				max, result[1] = array[k], array[k]
				max++
			}
		}
	}
	return result
}

//! 19

func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			if i-m[nums[i]] <= k {
				return true
			} 
		}
		m[nums[i]] = i
	}
	return false
}

//! 20

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
		if m[nums[i]] == true {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
//! 21


func intersectionisSubsequence(s string, t string) bool {
	count := 0
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(t); i++ {
		if s[count] == t[i] {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
}


//! 22 o(n)

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	m := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] {
			result = append(result, nums2[i])
			m[nums2[i]] = false
		}
	}

	return result
}

//! 23

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	m := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] >= 1 {
			result = append(result, nums2[i])
			m[nums2[i]]--
		}
	}
	return result
}


//! 24

func getCommon(nums1 []int, nums2 []int) int {
	m := make(map[int]bool)
    for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		fmt.Println(i)
		if m[nums2[i]] == true {
			return nums2[i]
		}
	}
	return -1
}

//! 25

func getCommon2(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] > nums2[j] { 
			j++
		} else {
			i++
		}
	} 
	return -1
}


//! 26

// func lengthOfLongestSubstring(s string) int {
// 	m := make(map[byte]bool)
// 	for i := 0; i < len(s); i++ {
// 		if m[s[i]] == true {
// 			return i
// 		}
// 		m[s[i]] = true
// 	}
// 	return -1
// }

//! 27x

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if m[t[i]] > 0 {
			m[t[i]]--
		} else {
			return t[i]
		}
	}
	return 0 
}


//! 28 

// func minSubArrayLen(target int, nums []int) int {
//     shrink := 0
// 	sum := 0
// 	count := 0
// 	for i := 0; i < len(nums); i++ {
// 		sum += nums[i]
// 		for sum >= target {
// 			sum -= nums[shrink]
// 			shrink++
// 			if sum == target  {
// 				if count == 0 {
// 					count = i - shrink
// 				} else if count > i - shrink {
// 					count = i-shrink
// 				}
// 			}
// 		}

// 	}
// 	return shrink
// }
//! 29 (leetcode 167) -> 0(n)/0(1) space complexity -> sorted

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right { 
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

//! 30 

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        m[s[i]]++
    }
    for i := 0; i < len(t); i++ {
        if m[t[i]] == 0 {
            return false
        }
        m[t[i]]--
    }
    return true
}


//! 31

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right  { 
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		} 
		if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		} 
		leftChar := unicode.ToLower(rune(s[left]))
		rightChar := unicode.ToLower(rune(s[right]))
		if leftChar != rightChar {
			return false
		}
		left++
		right--
		}
	return true
}
//! 32
func singleNumber(nums []int) int {
	result := 0
    for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

//! 33

func maxArea(height []int) int {

	maxArea := 0
    left, right := 0 , len(height)-1
	
	for left < right { 
		area := 0
		area = min(height[left], height[right]) * (right-left)
		if area > maxArea  {
			maxArea = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}

//! 34


func getConcatenation(nums []int) []int {
	a := len(nums)
    for i := 0; i < a; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}


//! 35 



// func lengthOfLongestSubstring(s string) int {
// 	longest := 0
// 	left := 0
// 	m := make(map[byte]bool)
// 	for i := 0; i < len(s); i++ {
// 		if  m[s[i]] == false {
// 			left++
// 			w := (i - left)+1
// 			longest = max(longest, w)
// 		}
// 		m[s[left]] = true
// 	} 
		
// 	return longest
// }


//! 36


func judgeCircle(moves string) bool {
	x := 0
	y := 0

    for i := 0; i < len(moves); i++ {
		switch string(moves[i]) {
			case "R":
				x++
			case "L":
				x--
			case "U":
				y++
			case "D":
				y--
		}
	}
	return x == 0 && y == 0 
}

//! 37

func findMaxConsecutiveOnes(nums []int) int {
    count := 0
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if maxCount < count {
				maxCount = count
			}
		} else{
			count = 0
		}
	}
	return maxCount
}

//! 38

func shuffle(nums []int, n int) []int {
    result := make([]int,0, 2*n)
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}
	return result
}

//! 39

func isMonotonic(nums []int) bool {
	count := 0
	first, last := 0, len(nums)-1

    for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] && nums[first] <= nums[last] {
			count++
		} else if nums[i] >= nums[i+1] && nums[first] > nums[last] {
			count--
		}
	}
	if -(count) == len(nums)-1{
		return true
	} else if count == len(nums)-1 {
		return true
	}
	return false
}
//! 40 -> 

func isMonotonic2(nums []int) bool {
    increasing := true
    decreasing := true


    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[i-1] {
            increasing = false
        }


        if nums[i] > nums[i-1] {
            decreasing = false
        }
    }

	fmt.Println(increasing)
	fmt.Println(decreasing)
    return increasing || decreasing
}

//! 41

func prefixSum(nums []int) []int {
	result := []int{}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result = append(result, sum)
	}
	return result		
}

//! 42

func fixedSize(nums []int, k int) []int {
	 shrink := 0
	 sum := 0
	 result := []int{}

	 for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i-shrink+1 == k {
			result = append(result, sum)
			sum -= nums[shrink]
			shrink++
		}
	 }
	 return result 
}

//! 43 -> ++++++++++++++++++++++++

func findMaxAverage(nums []int, k int) float64 {
    shrink := 0
	sum := 0
	maxResult := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i-shrink+1 == k {
			if maxResult == 0 {
				maxResult = sum
			} else if maxResult < sum {
				maxResult = sum
			}
			sum -= nums[shrink]
			shrink++
		}
	}
	return float64(maxResult)/float64(k)
}


func main() {


	nums:= []int{1,12,-5,-6,50,3}
	k := 4
	result := findMaxAverage(nums, k)
	fmt.Println(result)



	// nums:= []int{1,2,3,4,5,6}
	// k := 3
	// result := fixedSize(nums, k)
	// fmt.Println(result)



	// nums:= []int{1,2,3,4,5}
	// result := prefixSum(nums)
	// fmt.Println(result)

	// nums:= []int{1,2,3,4,5}
	// result := isMonotonic2(nums)
	// fmt.Println(result)


	// nums:= []int{4,4,4,4}
	// result := isMonotonic(nums)
	// fmt.Println(result)

	// nums:= []int{2,5,1,3,4,7}
	// n := 3
	// result := shuffle(nums, n)
	// fmt.Println(result)


	// nums:= []int{1,1,0,1,1,1}
	// result := findMaxConsecutiveOnes(nums)
	// fmt.Println(result)

	
	// moves:= "RRLLDDDUUU"
	// result := judgeCircle(moves)
	// fmt.Println(result)


	// word:= "abcabcbb"
	// result := lengthOfLongestSubstring(word)
	// fmt.Println(result)


	// nums:= []int{1,2,3,4}
	// result := getConcatenation(nums)
	// fmt.Println(result)

	// nums:= []int{1,8,6,2,5,4,8,3,7}
	// result := maxArea(nums)
	// fmt.Println(result)




	// fmt.Println(7^3)
	// nums:= []int{4,1,2,1,2,4,7}
	// result := singleNumber(nums)
	// fmt.Println(result)


	// s:= "a77a"
	// result := isPalindrome(s)
	// fmt.Println(result)


	// s:= "ab"
	// t := "aba"
	// result := isAnagram(s, t)
	// fmt.Println(result)


	// nums:= []int{1,2,3,4}
	// target := 6
	// result := twoSum(nums, target)
	// fmt.Println(result)



	// nums:= []int{2,3,1,2,4,3}
	// target := 7
	// result := minSubArrayLen(target, nums)
	// fmt.Println(result)



	// word1:= "a"
	// word2:= "aa"
	// result := findTheDifference(word1, word2)
	// fmt.Println(result)





	// word:= "pwwkew"
	// result := lengthOfLongestSubstring(word)
	// fmt.Println(result)


	// nums1 := []int{1,2,3,4,5,6,7,1,1,1,1,1,1,1,1}
	// nums2 := []int{4}
	// result := getCommon2(nums1, nums2)
	// fmt.Println(result)


	// nums1 := []int{19}
	// nums2 := []int{4,1,1,1,1,1,1,1,1,1,1,1,1,1,1,11}
	// result := getCommon(nums1, nums2)
	// fmt.Println(result)


	// nums1 := []int{4,9,5,4,1}
	// nums2 := []int{4,9,5,4,9,5,1,1}
	// result := intersect(nums1, nums2)
	// fmt.Println(result)



	// nums1 := []int{4,9,5}
	// nums2 := []int{4,9,5,4,9,5}
	// result := intersection(nums1, nums2)
	// fmt.Println(result)


	// word1 := ""
	// word2 := "ahbgdc"
	// result := isSubsequence(word1, word2)
	// fmt.Println(result)

	// nums := []int{1,2,3,4}
	// result := containsDuplicate(nums)
	// fmt.Println(result)


	// nums := []int{1,1}
	// target := 2
	// result := containsNearbyDuplicate(nums, target)
	// fmt.Println(result)



	// nums := []int{4, 2, 1, 3, 6}
	// result := LargestRange(nums)
	// fmt.Println(result)




	// word := "a,b,c,a,b"
	// result := FirstNonRepeatingCharacter2(word)
	// fmt.Println(result)



	// word := "aaaaaaaaaaaaaaaaaaaabbbbbbbbbbcccccccccccdddddddddddeeeeeeeeffghgh"
	// result := FirstNonRepeatingCharacter(word)
	// fmt.Println(result)


	// nums := []int{-5, -5, 2, 3, -2}
	// resultNums := ZeroSumSubarray(nums)
	// fmt.Println(resultNums)


	// nums := []int{2, 1, 5, 2, 3, 3, 4}
	// resultNums := FirstDuplicateValue(nums)
	// fmt.Println(resultNums)



	// nums := []int{2, 1, 5, 2, 3, 3, 4}
	// resultNums := sortingElement(nums)
	// fmt.Println(resultNums)

	// water := []int{1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1}
	// waterResult := MaxArea(water)
	// fmt.Println(waterResult)



	// array := []int{1, 2, 4, 5, 6}
	// targetElement := 3
	// resulstElement := MoveElementToEnd3(array, targetElement)
	// fmt.Println(resulstElement)


	// moveZero := []int{1,0,0,3,12}
	// resultZero := MoveZeroes2(moveZero)
	// fmt.Println(resultZero)

	// duplicates := []int{0,0,0,0}
	// resultDuplicates := removeDuplicates(duplicates)
	// fmt.Print(resultDuplicates,"\n")


	// seats := []int{1,1,1}
	// resultSeats := BestSeat(seats)
	// fmt.Println(resultSeats)



	// moton := []int{1, 2, 3, 3, 2, 1}
	// resultMoton := IsMonotonic(moton)
	// fmt.Println(resultMoton)


	// numUnsort := []int{-4,-23,0,3,10}
	// resultSort := SortedSquares(numUnsort)
	// fmt.Println(resultSort)



	// array := []int{5, 5, 5, 5, 5, 5, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12}
	// targetElement := 5
	// resulstElement := MoveElementToEnd(array, targetElement)
	// fmt.Println(resulstElement)

	// nums := []int{1,2,3,4,5,6,7}
	// target := 3
	// result := Rotate(nums, target)
	// fmt.Println(result)

}