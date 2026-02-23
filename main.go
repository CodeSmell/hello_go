// types of packages in Go
// executable packages: compiled and run as standalone programs
// library packages: provide reusable code that can be imported
package main

import (
	"fmt"
	"math/rand"
)

// a function that takes two strings
// and returns them in swapped order
// showing that Go supports multiple return values
func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	// declare variable
	var a string
	// assign value to variable
	a = "Hello"
	// declare and assign in one line
	b := "World"

	// swap on even
	if rand.Intn(10) % 2 == 0 {
		a, b = swap(a,b)
	}
			
	fmt.Println(a, b)
}