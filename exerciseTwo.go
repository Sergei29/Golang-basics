package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Create a program that reads the contents of a text file then prints its contents to the terminal.

The file to open should be provided as an argument to the program when it is executed at the terminal.  For example, 'go run main.go myfile.txt' should open up the myfile.txt file

To read in the arguments provided to a program, you can reference the variable 'os.Args', which is a slice of type string

To open a file, check out the documentation for the 'Open' function in the 'os' package - https://golang.org/pkg/os/#Open

What interfaces does the 'File' type implement?

If the 'File' type implements the 'Reader' interface, you might be able to reuse that io.Copy function!
*/
func exerciseTwo() {
	args := os.Args
	pathToFile := args[len(args)-1]

	if !isValidPathToFile(pathToFile) {
		fmt.Println(".txt file is missing as a last argument. please provide myFile.txt")
		os.Exit(1)
	}

	file, err := os.Open(pathToFile)

	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
		os.Exit(1)
	}

	outputTextFileV2(file)
}

func outputTextFileV2(file *os.File) {
	io.Copy(os.Stdout, file)
}

func outputTextFile(file *os.File) {
	bs := make([]byte, 10_000)
	length, errRead := file.Read(bs)

	if errRead != nil && errRead != io.EOF {
		fmt.Println("Error reading file: ", errRead.Error())
		os.Exit(1)
	}

	fmt.Println("File bytes length: ", length)

	fmt.Println(string(bs))
}

func isValidPathToFile(p string) bool {
	subStrings := strings.Split(p, ".")
	extension := subStrings[len(subStrings)-1]

	return extension == "txt"
}
