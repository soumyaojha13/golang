package main

import "fmt"

type Status int

const (
	Active Status = iota
	Inactive
	Pending
)

func main() {
	var s Status = Active

	fmt.Println(s) // 0
}
