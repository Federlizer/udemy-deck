package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

var (
	cardSuits  = [4]string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues = [13]string{"Ace", "Two", "Three", "Four",
		"Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}
)

func newDeck() deck {
	newDeck := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}
	return newDeck
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) String() string {
	return strings.Join(d, "\n")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.String()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(data), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		//find new random position
		newPosition := r.Intn(len(d) - 1)

		// swap the two elements
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
