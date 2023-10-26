package main

import (
	"fmt"
)

// START MAIN OMIT

func newCustomErr(msg string) error {
	return &customError{msg}
}

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func main() {
	var err error

	err = &customError{"custom error message 1"}
	fmt.Println(err.Error())

	err = newCustomErr("custom error message 2")
	fmt.Println(err.Error())
}

// END MAIN OMIT
