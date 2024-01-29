package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
	cray()

}

func newCard() string {
	return "5 of Diamonds"
}
