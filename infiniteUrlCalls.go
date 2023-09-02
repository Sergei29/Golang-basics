package main

import (
	"fmt"
	http "net/http"
	"time"
)

func infiniteUrlCalls(links []string) {
	ch := make(chan string)

	for _, link := range links {
		go checkUrl(link, ch)
	}

	/*
	 for a value yielded by channel `ch`, assign it to the variable `l`
	*/
	for l := range ch {
		go func(url string) {
			time.Sleep(time.Second * 5)
			checkUrl(url, ch)
		}(l)

		/*
			** Important: above, we are passing the `l` value as an argument
			to the `url string` parameter, because if we pass direct the
			`checkUrl(l, ch)` for example, wit hthe timer delay of 5s, possibly other calls will settle and the `ch` channel will yield to another value
		*/
	}
}

func checkUrl(url string, ch chan string) {
	_, err := http.Get(url)

	var msg string

	if err == nil {
		msg = url + " - is OK"
	} else {
		msg = "Error: " + err.Error()
	}

	fmt.Println(msg)
	ch <- url
}
