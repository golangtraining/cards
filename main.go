package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// ** deal
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println(hand.toString())
	// ** saveToFile
	// remainingDeck.saveToFile("my_cards")

	// ** newDeckFromFile
	// cards := newDeckFromFile("my")
	// cards.print()
}
