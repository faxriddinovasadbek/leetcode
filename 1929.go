package main

import(
	"fmt"
)

func getConcatenation(num []int) []int{
	num = append(num, num...)
	return num
}

func main(){
	fmt.Println(getConcatenation([]int{1,2,1}))
}