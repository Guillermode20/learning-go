package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	var num1, num2 int

	fmt.Println("Enter number one:")
	fmt.Scan(&num1)
	fmt.Println("Enter number two:")
	fmt.Scan(&num2)

	mult(num1, num2)
}

func mult(x int, y int) {
	fmt.Println("Your number:")
	fmt.Println(x*y)
}