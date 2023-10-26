package main

import (
	"fmt"
)

// START MAIN OMIT
// a zero-memory assignment that ensures Dog implements Animal
// even if it is never actually used as an Animal in the program
var _ Animal = (*Dog)(nil)

type Dog struct{}

func (d *Dog) Speak() { fmt.Println("dog says woof") }
func (d *Dog) Lick()  { fmt.Println("dog licks") }

type Animal interface {
	Speak()
	Move()
}

func main() {
	// now the compiler will complain about the Dog not implementing Animal
	// even though we don't use it as an Animal in the program
}

// END MAIN OMIT
