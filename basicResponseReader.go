package main

import (
	"fmt"
	"io"
	http "net/http"
	"os"
)

func basicResponseReader() {
	res, error := http.Get("http://www.google.com/")

	if error != nil {
		message := error.Error()
		fmt.Println("Error: ", message)
		os.Exit(1)
	}
	bs := make([]byte, 99_999)
	_, err := res.Body.Read(bs)

	if err != nil && err != io.EOF {
		message := err.Error()
		fmt.Println("Read Body Error: ", message)
		os.Exit(1)
	}

	html := string(bs)
	fmt.Println(html)
}
