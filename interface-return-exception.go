package main

import (
	"fmt"
	"time"
)

type Greeter interface{ SayHello(name string) }
type Leaver interface{ SayGoodbye(name string) }
type GreeterLeaver interface {
	Greeter
	Leaver
}

// START GETGREETER OMIT
type HappyGreeter struct{}

func (g *HappyGreeter) SayHello(name string)   { fmt.Println("Hello", name) }
func (g *HappyGreeter) SayGoodbye(name string) { fmt.Println("Goodbye", name) }
func (g *HappyGreeter) TellTime()              { fmt.Println(time.Now()) }

type GrumpyGreeter struct{}

func (g *GrumpyGreeter) SayHello(name string)   { fmt.Println("Go away", name) }
func (g *GrumpyGreeter) SayGoodbye(name string) { fmt.Println("Good riddance", name) }

// END GETGREETER OMIT

// START MAIN OMIT

// Let the function choose a specific implementation out of a set of
// possible implementations and return an interface
func NewGreeter(relationship string) GreeterLeaver {
	if relationship == "friend" {
		return &HappyGreeter{}
	}
	return &GrumpyGreeter{}
}

func GreetAndLeave(g GreeterLeaver, name string) {
	g.SayHello(name)
	g.SayGoodbye(name)
}

func main() {
	GreetAndLeave(NewGreeter("friend"), "Janet")
	GreetAndLeave(NewGreeter("enemy"), "Bob")
}

// END MAIN OMIT
