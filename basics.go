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

func main() {
	dog := &Dog{}
	dog.Speak()
	dog.Lick()

	cat := &Cat{}
	cat.Speak()
	cat.Scratch()
}

// END MAIN OMIT
