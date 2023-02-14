package main

func main() {
	// quan bich
	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
