package main

import "fmt"

func switchCase() {
	// switch expression
	// case value1:
	// code to be executed
	// default:
	// code to be executed if expression does not match any value
	fruit := "pinapple"
	switch fruit {
	case "apple":
		fmt.Println("It is an apple")
	default:
		fmt.Println("Unknown fruit")
	}
	// Multiple Conditions
	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday.")
	// case "Sunday":
	// 	fmt.Println("It's a weekend.")
	// default:
	// 	fmt.Println("Invalid day.")
	// }
}

func checkType(x interface{}) {}
