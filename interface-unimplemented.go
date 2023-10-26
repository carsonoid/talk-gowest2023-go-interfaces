package main

import (
	"fmt"
)

// START MAIN OMIT
type Dog struct{}

func (d *Dog) Speak() { fmt.Println("dog says woof") }
func (d *Dog) Lick()  { fmt.Println("dog licks") }

type Animal interface {
	Speak()
	Move()
}

func main() {
	// notice that we don't use the Dog as an Animal
	// even though we *intend* to use it as an Animal at some point
	// By default, the compiler will not complain about this
	// because it doesn't know that we intend to use it as an Animal
	d := &Dog{}
	d.Lick()
}

// END MAIN OMIT
