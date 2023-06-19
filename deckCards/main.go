package main

import (
	deck "udemy-projects.com/deck-cards/Deck"
)

func main() {
	cards := deck.NewDeck()

	hand, remainingCards := cards.Deal(4)
	hand.Print()
	remainingCards.Print()
}
