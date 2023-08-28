package main

import (
	"bufio"
	"fmt"
	"os"
)

// Constant vars: can't be updated
// If var name start with CAPITAL => Can be accessed PUBLICLY
const TokenVal string = "a12sje4d57"

func variables() {
	var username string = "Keshav"
	fmt.Println(username)
	fmt.Printf("Var type: %T \n", username) // Print type of variable

	// 8 bit Unsigned Int can store upto 255 max
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Var type: %T \n", smallVal)

	// Others: int8, uint32, int32 - to save memory
	// for Simplicity, just use: int

	// Variables decleared but not given any value
	// are by default initialized to 0.

	// Declaring vars
	// We can ignore keywords : var OR variableType while declaration
	// Lexar [Kind of grammer checker in GO] takes care of it.

	var website = "golang.in"
	fmt.Println(website)

	// NOTE: this is commonly used syntax
	// BUT: this is not allowed for Global Variables i.e out side a method/function.
	userCount := 3000
	fmt.Println(userCount)
}

func userInput() {
	// Reader reads the input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Rating: ")

	// Store the value read by Reader - Read till "\n" i.e "END of LINE"
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating. | ", input)

	// Note: "_" is for error handling
}

// Starting Point
func main() {
	variables()
	userInput()
}

/*
As in javascript almost everything is like object,
in Golang everything is a "type"
- strings, bool , int, float, complex etc.
*/
