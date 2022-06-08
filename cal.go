// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	num1 := 30
	num2 := 50
	result := 0
	operation := "+"
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
