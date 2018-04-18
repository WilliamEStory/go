package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	fmt.Println("Remaining")
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
