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



func main() {

	array := []int{5, 5, 5, 5, 5, 5, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12}
	targetElement := 5
	resulstElement := MoveElementToEnd(array, targetElement)
	fmt.Println(resulstElement)

	// nums := []int{1,2,3,4,5,6,7}
	// target := 3
	// result := Rotate(nums, target)
	// fmt.Println(result)

}