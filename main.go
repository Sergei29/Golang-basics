package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	cards = remainingCards
	cards.print()
}
