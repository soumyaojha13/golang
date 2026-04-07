package main

import "fmt"

func update(val *int) {
	*val = 50
}

func main() {
	x := 10
	update(&x)

	fmt.Println(x)
}
