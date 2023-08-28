package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func conversion() {
	fmt.Println("Enter Rating: ")

	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating. | ", rating)

	// "rating" type = string

	// convert string to float -
	// Since the string "rating" also contains "\n" @ end , we must trim it
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Adding 1 to rating: ", numRating+1)
	}

}

func main() {
	conversion()
}
