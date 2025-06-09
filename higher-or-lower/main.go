// Higher or Lower Game

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Higher or Lower!")

	number := rand.Intn(100)
	fmt.Println("Random Number 1-100 Generated!")
	
	guessed := false
	var playerGuess int

	for guessed != true {
		fmt.Println("Enter your choice!")
		fmt.Scanln(&playerGuess)

		if playerGuess > number {
			fmt.Println("Too high!")
		} else if playerGuess < number {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Correct! The number was", number)
			guessed = true
		}
	}
}
