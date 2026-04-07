package main

import (
	"fmt"
	"time"
)

// struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// method to change status
func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
}

// method to get amount
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	// instance 1
	order1 := order{
		id:     "1",
		amount: 50.0,
		status: "received",
	}
	order1.createdAt = time.Now()

	// instance 2
	order2 := order{
		id:     "2",
		amount: 120.5,
		status: "pending",
	}
	order2.createdAt = time.Now()

	// instance 3
	order3 := order{
		id:     "3",
		amount: 300.0,
		status: "shipped",
	}
	order3.createdAt = time.Now()

	// change status
	order1.changeStatus("completed")
	order2.changeStatus("cancelled")

	// print results
	fmt.Println(order1)
	fmt.Println(order2)
	fmt.Println(order3)

	// get amount
	fmt.Println("Order1 Amount:", order1.getAmount())
	fmt.Println("Order2 Amount:", order2.getAmount())
}

// constructor in go

func newPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}


//embeded strucks

package main

import "fmt"

// base struct
type Person struct {
    name string
    age  int
}

// embedded struct
type Employee struct {
    Person   // embedded struct
    id       int
    position string
}

func main() {
    emp := Employee{
        Person: Person{
            name: "Soumya",
            age:  22,
        },
        id:       101,
        position: "Developer",
    }

    fmt.Println(emp.name)     // accessed directly
    fmt.Println(emp.age)
    fmt.Println(emp.id)
}
