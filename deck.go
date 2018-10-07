package main

import (
	"io/ioutil"
	"strings"
)

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

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) String() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.String()), 0666)
}
