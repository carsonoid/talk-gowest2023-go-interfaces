package main

import (
	"fmt"
)

// START MAIN OMIT

type Dog struct{ age int }

func (d *Dog) Speak()         { fmt.Println("dog says woof") }
func (d *Dog) GetAge() string { return fmt.Sprintf("%d years", d.age) }

type AgeGetter interface {
	GetAge() string
}

type Animal interface {
	AgeGetter
	Speak()
}

func animalActions(a Animal) {
	a.Speak()
	fmt.Println("Age: " + a.GetAge())
}

func main() {
	animalActions(&Dog{age: 5})
}

// END MAIN OMIT
