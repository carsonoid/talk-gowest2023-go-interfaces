package main

// START MAIN OMIT
import (
	"errors"
	"fmt"
)

func main() {
	var err error

	err = errors.New("error message 1")
	fmt.Println(err.Error())

	err = fmt.Errorf("error message 2")
	fmt.Println(err.Error())
}

// END MAIN OMIT

// START ERR OMIT
type error interface {
	Error() string
}

// END ERR OMIT
