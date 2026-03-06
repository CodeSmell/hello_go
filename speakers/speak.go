package animal

type Speaker interface {
	Speak() string
}

type Walker interface {
	Walk() string
}

type WalkerBase struct{}

func (w WalkerBase) Walk() string {
	return "walking!"
}

type Animal struct {
	Name  string
	Class string
}

// Dog embeds Animal (composition)
type Dog struct {
	Animal     // Embedded - Dog has all Animal fields
	WalkerBase // Embedded - Dog has a default Walk method
	Breed      string
}

func NewDog(name string, breed string) Dog {
	return Dog{
		Animal: Animal{Name: name, Class: "Mammal"},
		Breed:  breed,
	}
}

// Dog implements the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Dog implements the Walk method of Animal
// func (d Dog) Walk() string {
// 	return "walking!"
// }

// Cat embeds Animal (composition)
type Cat struct {
	Animal     // Embedded - Cat has all Animal fields
	WalkerBase // Embedded - Cat has a default Walk method
}

func NewCat(name string) Cat {
	return Cat{
		Animal: Animal{Name: name, Class: "Mammal"},
	}
}

// Cat implements the Speak method of Animal
func (c Cat) Speak() string {
	return "Meow!"
}

// Cat implements the Walk method of Animal
// func (c Cat) Walk() string {
// 	return "walking!"
// }

// Not an Animal but still Speaks
type Echo struct {
	message string
}

func NewEcho(message string) Echo {
	return Echo{message: message}
}

func (e Echo) Speak() string {
	return e.message
}
