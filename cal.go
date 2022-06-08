// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var num1 int 
	var num2 int 
	var operation string
	fmt.Println("Enter First Number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second Number: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter operation")
	fmt.Scanln(&operation)

	var result int 

	if operation == "+" {
		result = num1 + num2
	}
	if operation == "-" {
		result = num1 - num2
	}
	if operation == "/" {
		result = num1 / num2
	}
	if operation == "*" {
		result = num1 * num2
	}
	fmt.Println("The result is: ", result)
}
