package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// chia bo bai ra lam 2 phan
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// 0666: read + write
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// get new deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		// Option 1 - log the error and return a call to newDeck()
		// option 2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// []byte -> string -> []string -> deck
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// create new rand
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// shuffer
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = swap(d[i], d[newPosition])
	}
}

func swap(a string, b string) (string, string) {
	return b, a
}
