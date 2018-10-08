package main

import (
	"fmt"
)

func main() {
	cards := newDeckFromFile("./my_deck.txt")
	fmt.Println(cards)
}
