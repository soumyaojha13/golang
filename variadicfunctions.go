//varidic function

package main

import "fmt"

func sum(nums ...int) {
    fmt.Println(nums)
}

func main() {
    sum(1, 2, 3)
}

// sum example

package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total
}

func main() {
    result := sum(1, 2, 3, 4)
    fmt.Println(result)
}

//slice passing

nums := []int{1, 2, 3}

sum(nums...)   // important