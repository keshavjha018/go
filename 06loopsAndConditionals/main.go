package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("----------- If/Else ----------")

	score := 46
	var res string

	if score > 45 {
		res = "Topper"
	} else if (score > 30) && (score <= 45) {
		res = "Average"
	} else {
		res = "Ok"
	}

	fmt.Println("Student is", res)

	if score%2 == 0 {
		fmt.Println("Score is even")
	}

	// --------------------------------------------
	fmt.Println("----------- Switch:Case ----------")

	rand.Seed(time.Now().UnixNano())
	randomVal := rand.Intn(51) // random no bw 0 : 50

	switch randomVal {
	case 50:
		fmt.Println("Allowed to Pass")
		fallthrough // exceptionally allows to check for other cases

	default:
		fmt.Println("Not Allowed")

	}

	// --------------------------------------------
	fmt.Println("----------- Loops ---------------")

	dates := []int{1, 2, 3, 4}
	fmt.Println("Date: ", dates)

	// for i := 0; i < len(dates); i++ {
	// 	fmt.Printf("%v ", dates[i])
	// }

	// remember: here "i" gives index not val here (unlike c++)
	// for i := range dates {
	// 	fmt.Printf("%v ", dates[i])
	// }

	// Directly access values
	for index, date := range dates {
		fmt.Println("At ", index, ", date: ", date)
		// break
		// continue
	}
}
