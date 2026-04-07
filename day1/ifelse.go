package main

import "fmt"

func main() {

	age := 16

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age > 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a kid")
	}

}
