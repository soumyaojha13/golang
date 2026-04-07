package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("error occured")

	}

	defer file.Close()

	buf := make([]byte, 10)

	d, err := file.Read(buf)

	if err != nil {
		fmt.Println("error occured")

	}

	fmt.Println("data", d, buf)

}
