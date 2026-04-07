package main

import (
	"fmt"
	"time"
)

func processnum(numchan chan int) {
	fmt.Println("processing  number", <-numchan)
}

func main() {
	numchan := make(chan int)

	go processnum(numchan)

	numchan <- 5

	time.Sleep(time.Second * 2)
}
