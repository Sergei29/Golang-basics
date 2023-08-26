package main

import (
	"fmt"
	"io"
	http "net/http"
	"os"
)

func advancedResponseReader() {
	res, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if error != nil {
		message := error.Error()
		fmt.Println("Error: ", message)
		os.Exit(1)
	}

	_, err := io.Copy(os.Stdout, res.Body)

	if err != nil && err != io.EOF {
		fmt.Println("Output Error: ", err.Error())
		os.Exit(1)
	}
}
