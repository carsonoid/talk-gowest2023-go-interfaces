package main

import "time"

// START MAIN OMIT
type Greeter interface{ SayHello(name string) }
type Leaver interface{ SayGoodbye(name string) }
type GreeterLeaver interface {
	Greeter
	Leaver
}
type Client interface {
	GreeterLeaver
	TellTime() time.Time
	GetTheAnswer() int
	DeliverTheAnswer(answer int)
	FetchWater() (int, error)
}

func Greet(g Greeter) { g.SayHello("Bob") }

func GreetAndLeave(l GreeterLeaver) {
	l.SayHello("Bob")
	l.SayGoodbye("Bob")
}

// END MAIN OMIT

func main() {

}
