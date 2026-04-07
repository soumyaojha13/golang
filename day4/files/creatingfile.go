package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("example2.txt")

	if err != nil {
		fmt.Println("error occured")

	}

	defer file.Close()

	file.WriteString("go lang")

}
