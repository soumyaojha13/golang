package main

import "fmt"

// declaring array

func main() {
	var nums [4]int

	// adding elements to array
	nums[0] = 1
	nums[2] = 2
	fmt.Println(len(nums))
	fmt.Println(nums)

	// adding elements while declaring array

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	// 2d array

	var arr [2][3]int
}
