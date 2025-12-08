package main

import "fmt"

func array() {
	// var arrayName [size]elementType
	var numbers [5]int
	fmt.Println(numbers)
	numbers[4] = 20
	fmt.Println(numbers)
	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "grapes"}
	fmt.Println("Fruits array", fruits)

	fmt.Println("Third element", fruits[2])
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[0] = 100
	fmt.Println(originalArray, copiedArray)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
	originalArray1 := [3]int{1, 2, 3}
	var copiedArray1 *[3]int
	copiedArray1 = &originalArray1
	copiedArray1[0] = 20032
	// it changed here as it point to the addresses
	fmt.Println(originalArray1, copiedArray1)
}
