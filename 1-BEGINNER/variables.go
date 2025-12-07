package main

import "fmt"

// global
var middleName = "viral"

func variables() {
	var age int
	var name string = "john"
	name1 := "Jane"
	fmt.Print(name, age, name1)
	// var name :=12 we dont use short variable decalration
	// default values
	// numeric type: 0
	// boolean type false
	// string type ""
	// pointers, slice, maps, functions and structs:nil
}

func printName() {
	fitstName := "wrg"
	fmt.Println(fitstName)
	// fitsName is destroyed as outside the function scope
	// variables are mutable
}
