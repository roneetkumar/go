package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Exepected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Exepected first card, Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs of Jack" {
		t.Errorf("Exepected first card, Clubs of Jack, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 169 {
		t.Errorf("Exepected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
