package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Excpectd length of 12,but got %v", len(d))
	}
	if d[0] != "Ace Of Spades" {
		t.Errorf("Excpectd first card of ace of spades,but got %v", d[0])
	}
	if d[len(d)-1] != "Three Of Clubs" {
		t.Errorf("Excpectd last card of three of clubs,but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktest")
	d := newDeck()
	d.saveToFile("_decktest")
	ldd := newDeckFromFile("_decktest")
	if len(ldd) != 12 {
		t.Errorf("Excpectd 12 cards,but got %v", len(ldd))
	}
	os.Remove("_decktest")
}
