package main

import (
	"fmt"
	"time"
)

// START MAIN OMIT
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

// END MAIN OMIT

func NewHappyGreeter() GreeterLeaver {
	return &HappyGreeter{}
}

func GreetAndLeave(g GreeterLeaver, name string) {
	g.SayHello(name)
	g.SayGoodbye(name)
}

func main() {
	hg := NewHappyGreeter()
	GreetAndLeave(hg, "Bob")
	hg.TellTime()
}

// END GETGREETER OMIT
type GrumpyGreeter struct{}

func (g *GrumpyGreeter) SayHello(name string)   { fmt.Println("Go away", name) }
func (g *GrumpyGreeter) SayGoodbye(name string) { fmt.Println("Good riddance", name) }
func (g *GrumpyGreeter) ChangeMood()            { /* ... */ }
