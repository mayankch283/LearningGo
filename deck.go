package main

import (
	"fmt"
	"os"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ", ")

}

func (d deck) saveToFile() error {
	return os.WriteFile("cards", []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) string {
	cards, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(cards)
}
