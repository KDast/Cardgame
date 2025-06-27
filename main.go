package main

import (
	"fmt"
)

func main() {
	cardDeck := makeCardDeck()
	for i, card := range cardDeck {
		n := i + 1
		fmt.Printf("card : %v is %v\n", n, card.name)
	}
	shuffledDeck := shuffle(cardDeck)
	for i, card := range shuffledDeck {
		n := i + 1
		fmt.Printf("card : %v is %v\n", n, card.name)
	}
}
