package main

import (
	"fmt"
	http "net/http"
)

func urlCalls(links []string) {
	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
}

func checkLink(url string, ch chan string) {
	_, err := http.Get(url)

	var msg string

	if err == nil {
		msg = url + " - is OK"
	} else {
		msg = "Error: " + err.Error()
	}

	ch <- msg
}
