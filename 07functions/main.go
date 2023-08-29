package main

import "fmt"

func adder(a int, b int) int {
	return a + b
}

// accept any no of arguments
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// return more than two vals
func multipleReturns(val int) (int, string) {
	return val, "Returning addtional String"
}

//--------------- DEFERS ----------------------
// executes at the end of function
// if more than 1 defers => executes in LIFO order

// Use Cases: Closing the file/db at function end in apps

func testDefer() {
	defer fmt.Println("World")
	defer fmt.Println("Earth")
	fmt.Println("Hello")

	// Output
	/*
		Hello
		Earth
		World
	*/
}

func main() {
	fmt.Println(adder(1, 3))
	fmt.Println(proAdder(2, 3, 4, 5, 6, 7))

	res, msg := multipleReturns(2)
	fmt.Println(res, msg)

	newUser := User{"Keshav", 23}

	// Methods
	newUser.getAge()

	// DEFERS
	testDefer()
}

// METHODS
type User struct {
	Name string
	age  int
}

func (u User) getAge() {
	fmt.Println("age: ", u.age)

	// also we can change vals: by  passing pointer/reference
	// u.age = u.age + 1
	// fmt.Println("New age: ", u.age)
}
