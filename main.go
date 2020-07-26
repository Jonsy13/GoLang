package main

import (
	"fmt"
)

// Calculate is used to calculate increment.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go Testing Hii")
	result := Calculate(2)
	fmt.Println(result)
}
