package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	divider()
	hand, newDeck := cards.deal(5)
	fmt.Println(hand, "\n", newDeck, "\n", len(newDeck))
	cards.print()
	divider()
}

func divider() {
	fmt.Println("****************************")
}
