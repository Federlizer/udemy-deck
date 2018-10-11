package main

import (
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}

	var spades int
	var diamonds int
	var hearts int
	var clubs int

	for _, card := range d {
		if strings.HasSuffix(card, "Spades") {
			spades++
		}
		if strings.HasSuffix(card, "Diamonds") {
			diamonds++
		}
		if strings.HasSuffix(card, "Hearts") {
			hearts++
		}
		if strings.HasSuffix(card, "Clubs") {
			clubs++
		}
	}
	if spades != diamonds && diamonds != hearts && clubs != spades {
		t.Errorf("Not the right cards in the deck")
	}
}
