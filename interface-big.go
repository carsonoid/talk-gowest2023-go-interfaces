package main

import "time"

// START MAIN OMIT
type Client interface {
	SayGoodbye(name string)
	SayHello(name string)
	TellTime() time.Time
	GetTheAnswer() int
	DeliverTheAnswer(answer int)
	FetchWater() (int, error)
}

func Greet(c Client) { c.SayHello("Bob") }

func GreetAndLeave(c Client) {
	c.SayHello("Bob")
	c.SayGoodbye("Bob")
}

// END MAIN OMIT

func main() {

}
