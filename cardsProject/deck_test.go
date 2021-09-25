package main

import (
	"testing"

)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck of length of 51, but returned %v", len(d))
	}
	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King Of Clubs" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[len(d)-1])
	}

}