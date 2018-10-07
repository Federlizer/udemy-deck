package main

import "fmt"

type deck []string

var (
	cardSuits  = [4]string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
)

func newDeck() deck {
	newDeck := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, suit+" of "+value)
		}
	}
	return newDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
