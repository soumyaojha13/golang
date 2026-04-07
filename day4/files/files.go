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
	fileinfo, err := file.Stat()

	if err != nil {
		fmt.Println("error occured")
	}

	fmt.Println("file name:", fileinfo.Name())
	fmt.Println("file size:", fileinfo.Size())

}
