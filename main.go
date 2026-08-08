package main

import "fmt"

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


func LeftRightDifference(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		left := 0
		right := 0

		for k := 0; k < len(nums); k++ {
			if k < i {
				left += nums[k]
			} else if k > i {
				right += nums[k]
			}
		}
		if left > right {
			result = append(result, -1)
		} else if left < right {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

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

//! 10



func main() {

	seats := []int{1,1,1}
	resultSeats := BestSeat(seats)
	fmt.Println(resultSeats)



	// moton := []int{1, 2, 3, 3, 2, 1}
	// resultMoton := IsMonotonic(moton)
	// fmt.Println(resultMoton)


	// numDifference := []int{1,2,3,4}
	// resultDifference := LeftRightDifference(numDifference)
	// fmt.Println(resultDifference)



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