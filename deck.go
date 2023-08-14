package main

import "fmt"

type Deck []string

var cardTitles [13]string = [13]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

var cardSuits [4]string = [4]string{"Hearts", "Diamonds", "Spades", "Clubs"}

func (d Deck) print() {
	for index, card := range d {
		fmt.Println(index+1, card)
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
