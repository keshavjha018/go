package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("---------- ARRAY -----------------")
	// ----------------------------------------------

	var arr [4]string // must mention size
	arr[0] = "Apple"
	arr[1] = "Mango"
	arr[3] = "Tomato"

	// Note: if arrr[2] is empty => prints empty space

	fmt.Println(arr)
	fmt.Println("length of array: ", len(arr))

	// declare array elements straight
	var vegs = [4]string{"eggplant", "potato"} // empty @ end 2
	fmt.Println(vegs)

	// -----------------------------------------------
	fmt.Println("----------  SLICES -----------------")
	// ------------------------------------------------

	// kind of vector i.e better version Arrays

	var fruits = []string{}
	fmt.Println(fruits)

	// add as many values as we want - append
	fruits = append(fruits, "Tomato", "Banana", "Apple")
	fmt.Println(fruits)

	// slice part of the array
	var topTwoFruits = append(fruits[:2])
	fmt.Println("Top 2 fruits are : ", topTwoFruits)

	// for append to use => data type must be "slice"
	// meaning size cant be declared while initialization

	// But in case of Dynamic allocation
	// even if size is declared: we can append

	scores := make([]int, 2)
	scores[0] = 23
	scores[1] = 45

	scores = append(scores, 34, 54)
	fmt.Println("scores: ", scores)

	// More Operations

	// Sort
	sort.Ints(scores)
	fmt.Println("SORTED: ", scores)
	fmt.Println("Is Sorted: ", sort.IntsAreSorted(scores))

	// remove ele
	var indexToRemove int = 2
	scores = append(scores[:indexToRemove], scores[indexToRemove+1:]...)
	fmt.Println("Deleted 2nd index: ", scores)

	// ----------------------------------------------
	fmt.Println("----------  HASHMAP -----------------")
	// ------------------------------------------------

	nameMap := make(map[int]string)
	nameMap[70] = "keshav"
	nameMap[33] = "Brij"

	fmt.Println("Map : ", nameMap)

	// traversing a map: loop
	for roll, name := range nameMap {
		fmt.Println("Roll: ", roll, ", Name: ", name)
	}

	// ----------------------------------------------
	fmt.Println("----------  STRUCTS -----------------")
	// ------------------------------------------------

	// there is no Inheritance , super or Parent concept in golang
	newUser := User{"Keshav", "keshav@email.com", true, 22}
	fmt.Println("New User Struct: ", newUser)
	fmt.Println("Age of User: ", newUser.age)

	// detailed info	: Whole Structure(key-val pairwise)
	fmt.Printf("New User Full Details: %+v", newUser)

}

// Any variable which starts with Capital Can be accessed from anywhere

type User struct {
	Name       string
	Email      string
	IsVerified bool
	age        int
}
