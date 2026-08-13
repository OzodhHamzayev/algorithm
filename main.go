package main

import (
	"fmt"
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



func main() {


	nums := []int{4, 2, 1, 3, 6}
	result := LargestRange(nums)
	fmt.Println(result)




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