package main

func main() {
	cards := newDeck() //var cards deck = newDeck()

	cards.shuffle()
	cards.print()
}
