package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of Strings
type deck []string

func newDeck() deck {
	cards := deck{}
	// create suits and values to add the card
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// add to the card
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
