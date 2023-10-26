package main

import "fmt"

// START MAIN OMIT
type Client struct{}

func (c Client) SayHello(name string)   { fmt.Println("Hello", name) }
func (c Client) SayGoodbye(name string) { fmt.Println("Goodbye", name) }

func main() {
	greetUser(Client{}, "Bob")
}

type greeter interface {
	SayHello(name string)
}

func greetUser(g greeter, name string) {
	g.SayHello(name)
}

// END MAIN OMIT
