package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("cards")
	cards := newDeckFromFile("cards")
	cards.shuffle()
	cards.print()
}
