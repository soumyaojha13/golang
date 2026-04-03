func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println(result)
}

//multiple return values

func calc(a int, b int) (int, int) {
    return a + b, a - b
}

func main() {
    sum, diff := calc(10, 5)
    fmt.Println(sum, diff)
}

// function as variable

package main

import "fmt"

func main() {

    add := func(a, b int) int {
        return a + b
    }

    fmt.Println(add(2, 3))
}

//functions passed as parameters

func operate(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    result := operate(5, 3, func(x, y int) int {
        return x * y
    })

    fmt.Println(result)
}