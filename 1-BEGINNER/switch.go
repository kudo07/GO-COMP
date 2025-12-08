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
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case int32:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}
