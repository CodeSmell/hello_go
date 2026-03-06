// types of packages in Go
// 1) executable packages: compiled and run as standalone programs
// 2) library packages: provide reusable code that can be imported
package main

// in Go it is typical to make the module path match the URL for the repository.
// but the module path can be whatever you want (and is found in the go.mod file)
import (
	"example/hello/morestrings"

	// standard library imports
	"fmt"
	"math/rand"

	// third party imports
	// use go mod tidy to add this to your go.mod file
	"github.com/google/go-cmp/cmp"
)

// a function that takes two strings
// and returns them in swapped order
// showing that Go supports multiple return values
func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	// declare variable
	// assign value to variable on separate line
	var a string
	a = "Hello"

	// declare and assign in one line
	b := "World"

	// declare and assign with type inference
	var goLang = "goLang"

	// swap on even
	if rand.Intn(10)%2 == 0 {
		a, b = swap(a, b)
	} else {
		b = goLang
	}

	fmt.Println(a, b)

	// demonstrate the use of the morestrings package
	fmt.Println("ReverseRunes:", morestrings.ReverseRunes(goLang))

	// demonstrate the use of the cmp package
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
