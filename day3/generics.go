//generics - used define a common datatype to a function instead defining a new function everytime for another function

package main

import "fmt"

func printslice[T any](items []T) { // instead of any interfaces{} can be used
	for _, items := range items {
		fmt.Println(items)
	}
}

func main() {
	//nums := []int{1, 2, 3, 4}
	language := []string{"hindi", "english"}
	printslice(language)
}
