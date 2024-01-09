package main

import "fmt"

func maximumCount(nums []int) int {
	musbat := 0
	manfiy := 0
	for _, num := range nums {
		if num > 0{
			musbat++
		}else{
			manfiy++
		}
	}

	if manfiy > musbat{
		return manfiy
	}else{
		return musbat
	}
}

func main() {
	fmt.Println(maximumCount([]int{-2,-1,-1,1,2,3}))
}
