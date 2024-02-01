package main

import (
	"fmt"
)

func main() {
	cards := newDeck() //var cards deck = newDeck()

	cards.saveToFile()
	fmt.Println(newDeckFromFile("nigga"))
	fmt.Println(newDeckFromFile("cards"))

}
