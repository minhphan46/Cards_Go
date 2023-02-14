package main

import "testing"

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
