package main

import (
	"fmt"
)

// START MAIN OMIT

type intErr int

func (i intErr) Error() string {
	return fmt.Sprintf("int error: %d", i)
}

func main() {
	var err error

	err = intErr(1)
	fmt.Println(err.Error())
}

// END MAIN OMIT
