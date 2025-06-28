package main

import (
	"fmt"
)

func main() {
	cardDeck := makeCardDeck()
	shuffledDeck := shuffle(cardDeck)
	var playerOne player
	var playerTwo player
	playerOne.name = "player 1"
	playerTwo.name = "player 2"

	playerOne.hand, playerTwo.hand = distributeCards(shuffledDeck)
	winner, _ := war(playerOne, playerTwo)
	fmt.Printf("THE WINNER IS:\n %v\n", winner.name)
}
