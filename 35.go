package main

import "fmt"

func searchInsert(nums []int, target int) int {
	number := 0

	check := true

	for _, num := range nums{
		if num == target{
			check = false
			break
		}
		number++
	}

	if check == true{
		number = 0
		for _, num := range nums{
			if num > target{
				break
			}
			number++
		}
	}
	return number
}

func main() {

	fmt.Println(searchInsert([]int{1,3,5,6}, 2))

}