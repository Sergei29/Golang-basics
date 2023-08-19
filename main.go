package main

var filename string = "assets/my_cards"

func main() {
	retrievedDeck := newDeckFromFile(filename)
	retrievedDeck.print()
	retrievedDeck.shuffle()
	retrievedDeck.print()
}
