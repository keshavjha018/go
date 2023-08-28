package main

import (
	"fmt"
	"time"
)

func handleTime() {
	fmt.Println("Time Operations !!")

	// Today's Date
	currTime := time.Now()
	fmt.Println(currTime)

	// Formatting: Rule [fixed string]
	fmt.Println(currTime.Format("01-02-2006 15:04:05 Monday"))

	// Define a date urself
	myDate := time.Date(2020, time.December, 31, 12, 12, 32, 0, time.UTC)
	fmt.Println(myDate)
}

func main() {
	handleTime()
}

/*

	BUILD (exe to share)
	GOOS="windows" go build

*/
