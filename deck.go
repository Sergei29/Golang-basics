package main

import (
	"fmt"
	"os"
	"strings"
)

type Deck []string

var cardTitles [13]string = [13]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

var cardSuits [4]string = [4]string{"Hearts", "Diamonds", "Spades", "Clubs"}

func (d Deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func newDeck() Deck {
	deck := Deck{}

	for _, suit := range cardSuits {
		for _, title := range cardTitles {
			deck = append(deck, title+" of "+suit)
		}
	}

	return deck
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func toByteSlice(s string) []byte {
	return []byte(s)
}

func (d Deck) saveToFile(fileName string) error {
	data := toByteSlice(d.toString())
	return os.WriteFile(fileName, data, 0666)
}

func byteSliceToDeck(b []byte) Deck {
	stringifiedBs := string(b)
	sliceOfCards := strings.Split(stringifiedBs, ",")

	return Deck(sliceOfCards)
}

func newDeckFromFile(fileName string) Deck {
	data, error := os.ReadFile(fileName)

	if error != nil {
		// error handling...
		// Option 1: log the error, and return a new Deck
		// Option 2: log the error, and entirely quit the program
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	return byteSliceToDeck(data)
}
