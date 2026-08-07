package main


import ("fmt")


func main() {
	 
	word := []string{"H","a","n","n","a","h"}
	result := ReverseString(word)
	fmt.Println(result)
}


func  ReverseString(s []string) []string {
	result := []string{}
	for i := len(s)-1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}
