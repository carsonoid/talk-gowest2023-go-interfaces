package main

import (
	"fmt"
)

// START MAIN OMIT

type Dog struct{ age int }

func (d *Dog) Speak() { fmt.Println("dog says woof") }
func (d *Dog) Lick()  { fmt.Println("dog licks") }

type Cat struct{ age int }

func (c *Cat) Speak()   { fmt.Println("cat says meow") }
func (c *Cat) Scratch() { fmt.Println("cat scratches") }

type Animal interface {
	Speak()
}

func main() {
	// Dog and Cat are different types, so we can't put them into a slice together
	// But they both implement the Animal interface so we can make a slice of Animals
	// and put them in there
	animals := []Animal{&Dog{}, &Cat{}}
	for _, animal := range animals {
		animal.Speak()
	}
}

// END MAIN OMIT
