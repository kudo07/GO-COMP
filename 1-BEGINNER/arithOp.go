package main

import (
	"fmt"
	"math"
)

func arithOp() {
	// the value reads from right to left reading rigth to left
	// overflow
	var i int8 = 100
	var j int8 = 29
	var result int8
	result = i + j
	// -127
	fmt.Println(result)
	var smallFloat float64 = 1.0e-323
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
	// loss of precision revert to 0
}
