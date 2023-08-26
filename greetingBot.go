package main

import . "fmt"

type Bot interface {
	getGreeting() string
}
type EnglishBot struct{}
type SpanishBot struct{}

func greetingBot() {
	engBot := EnglishBot{}
	spainBot := SpanishBot{}

	printGreeting(engBot)
	printGreeting(spainBot)
}

func printGreeting(b Bot) {
	Println(b.getGreeting())
}

func (eb EnglishBot) getGreeting() string {
	// some logic very specific to english greeting...
	return "Hi there!"
}

func (sb SpanishBot) getGreeting() string {
	// some logic very specific to spanish greeting...
	return "Hola!"
}
