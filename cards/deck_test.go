package main

import "testing"

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
