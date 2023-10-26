package main

import (
	"fmt"
	"io/fs"
	"os"
)

// START ERR OMIT
func newFileInfoErr(fileInfo fs.FileInfo) error {
	return &fileInfoErr{
		FileInfo: fileInfo,
	}
}

type fileInfoErr struct {
	fs.FileInfo
}

func (f fileInfoErr) Error() string {
	return fmt.Sprintf("file info error: %s, mode: %s, modtime: %s",
		f.Name(), f.Mode(), f.ModTime())
}

// END ERR OMIT

// START MAIN OMIT
func main() {
	fileInfo, err := os.Stat("README.md")
	if err != nil {
		panic(err)
	}
	err = newFileInfoErr(fileInfo)
	fmt.Println(err.Error())

	// you can also use type assertion to get the underlying type
	// back from the error if you need to access its fields
	if fileInfoErr, ok := err.(*fileInfoErr); ok {
		fmt.Println("-fileInfoErr Name-\n" + fileInfoErr.Name())
	}
}

// END MAIN OMIT
