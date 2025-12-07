package main

import "fmt"

func loops() {
	// simple iteration over a range
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	// iterate over collection
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index:%v", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd number", i)
		if i == 5 {
			break
		}
	}
	for i := range 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println("We have a lift off!")
}
