package main

import (
	"os"
	"testing"
)

// this file to test the file deck.go

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Excepted deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Excepted first card of Ace of Spadesgo, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Excepted last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// testing saveToDeck + newDeckFromFile
// delete any files in current working directory with the name "_decktesting"
// create a deck
// save to file
// load from file
// Asset deck length
// delete any files in current working directory with the name "_decktesting"

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
