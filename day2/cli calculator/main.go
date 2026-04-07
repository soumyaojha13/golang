package main

import "fmt"

func main() {
	var a, b float64
	var op string
	var repeat string = "yes"

	fmt.Println("CLI CALCULATOR")

	for repeat != "no" {
		// take input
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)

		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scan(&op)

		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		// perform calculation
		switch op {
		case "+":
			fmt.Println("Result:", a+b)
		case "-":
			fmt.Println("Result:", a-b)
		case "*":
			fmt.Println("Result:", a*b)
		case "/":
			if b != 0 {
				fmt.Println("Result:", a/b)
			} else {
				fmt.Println("Error: Division by zero")
			}
		default:
			fmt.Println("Invalid operator")
		}

		// ask user to continue
		fmt.Print("Continue? (yes/no): ")
		fmt.Scan(&repeat)
	}

	fmt.Println("Calculator exited.")
}
