//shape interface with circle and rectangle structs

package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape

	s = Rectangle{width: 10, height: 5}
	fmt.Println("Rectangle Area:", s.area())

	s = Circle{radius: 7}
	fmt.Println("Circle Area:", s.area())
}
