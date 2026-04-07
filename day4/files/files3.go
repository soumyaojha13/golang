package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println("error occured")

	}

	fmt.Println(string(file))

}
