package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("Expect deck length of 52, but got: %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first car of Ace of Spades, but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Ten of Clubs" {
		t.Errorf("Expected first car of Ten of Clubs, but got %v", cards[len(cards)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	_ = os.Remove("_deckTesting")

	d := newDeck()
	_ = d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck but got %v", len(loadedDeck))
	}

	_ = os.Remove("_deckTesting")
}
