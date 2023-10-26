package main

import (
	"fmt"
)

// START MAIN OMIT

type Dog struct{ age int }

func (d *Dog) Speak()         { fmt.Println("dog says woof") }
func (d *Dog) Lick()          { fmt.Println("dog licks") }
func (d *Dog) GetAge() string { return fmt.Sprintf("%d years", d.age) }

type Animal interface{ Speak() }
type Licker interface{ Lick() }
type AgeGetter interface{ GetAge() string }

func animalSpeak(a Animal)     { a.Speak() }
func lickerLick(l Licker)      { l.Lick() }
func ageGetterSay(a AgeGetter) { fmt.Printf("is %s old\n", a.GetAge()) }

func main() {
	d := &Dog{age: 4}

	// dog satisfies all 3 interfaces so it can be passed to all functions
	// you do not need to declare implemented interfaces
	animalSpeak(d)
	lickerLick(d)
	ageGetterSay(d)
}

// END MAIN OMIT
