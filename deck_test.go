package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	nd := newDeck()

	if len(nd) != 16 {
		t.Errorf("Expected deck of length 16, but got %v", len(nd))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") //returns err which is yet to be handled

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck of length 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
