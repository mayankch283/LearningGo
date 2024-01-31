package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := []string{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"} //Enough elements to test. Should eventually go up to King.

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
