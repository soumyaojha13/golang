//slice - dyanmic (array)
//it has no fixed size

package main

import "fmt"

func main() {
	//empty slice
	var nums []int
	fmt.Println(nums)
	//declare with values
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(nums1)

	//using make()
	nums2 := make([]int, 2, 5)

	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	fmt.Println(nums2)

}

// slice operator

arr := [5]int{1, 2, 3, 4, 5}

slice := arr[1:4]

//slice package

contins different fumction

// 2d slices

slice := [][]int{
	{1, 2, 3},
	{4, 5, 6},
}

fmt.Println(slice)