package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected to have 52 cards, but got %v", len(deck))
		return
	}

	if deck[0] != "Two of Hearts" {
		t.Errorf("Expected to have 'Two of Hearts' as the first card, but got %v", deck[0])
	}

	if deck[51] != "Ace of Clubs" {
		t.Errorf("Expected to have 'Ace of Clubs' as the last card, but got %v", deck[51])
	}
}
