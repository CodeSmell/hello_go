// types of packages in Go
// 1) executable packages: compiled and run as standalone programs
// 2) library packages: provide reusable code that can be imported
package main

// in Go it is typical to make the module path match the URL for the repository.
// but the module path can be whatever you want (and is found in the go.mod file)
import (
	// local package imports
	"example/hello/distance"
	"example/hello/morestrings"
	speakers "example/hello/speakers"

	// standard library imports
	"fmt"
	"log"
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

// initial playground code to
// work with Go syntax
func hello_sandbox() {
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

	_, err := morestrings.ReverseRunesWithErr("")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func distance_sandbox() {
	dist1 := distance.Kilometers(100.0)
	dist2 := dist1.ToMiles()

	fmt.Println(dist1, "in kilometers is", dist2, "in miles")
	// compiler error due to mismatched types
	//result := dist1 + dist2
	//fmt.Println(result)
}

// a function that takes anything that can speak
func talk_to_me(speaker speakers.Speaker) {
	fmt.Println(speaker.Speak())
}

// a function that takes anything that can move
func walk_this_way(walker speakers.Walker) {
	fmt.Println(walker.Walk())
}

func main() {
	fido := speakers.NewDog("Fido", "Labrador")
	talk_to_me(fido)
	walk_this_way(fido)

	garfield := speakers.NewCat("Garfield")
	talk_to_me(garfield)
	walk_this_way(garfield)

	echo := speakers.NewEcho("Hellooo...")
	talk_to_me(echo)
	// echo does not walk
	// walk_this_way(echo)
}
