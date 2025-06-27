package main

import (
	"log"
	"math/rand"
	"strconv"
)

type cards struct {
	name     string
	numberID int
	suit     string
	value    int
}

func makeCardDeck() []cards {
	var cardDeck []cards
	cardSuits := [4]string{"heart", "diamond", "spade", "club"}
	for p, suits := range cardSuits {
		n := 13 * (p)
		for i := 1; i < 14; i++ {
			var cardName string
			if i == 1 {
				cardName = "ace" + " " + suits
			} else if i == 11 {
				cardName = "jack" + " " + suits
			} else if i == 12 {
				cardName = "queen" + " " + suits
			} else if i == 13 {
				cardName = "King" + " " + suits
			} else {
				number := strconv.Itoa(i)
				cardName = number + " " + suits
			}

			var card = cards{
				name:     cardName,
				numberID: i + n,
				suit:     suits,
				value:    i,
			}
			cardDeck = append(cardDeck, card)
		}
	}
	return cardDeck
}

func shuffle(deck []cards) []cards {
	var posList []int
	var shuffledCards []cards
	for i := range 52 {
		posList = append(posList, i)
	}

	for {
		if len(deck) == 0 || len(posList) == 0 {
			break
		}
		x := rand.Intn(len(posList))
		if x > len(posList) {
			log.Fatal("x is out of range")
		}
		cardPosition := posList[x]
		for p := range posList {
			if posList[p] == cardPosition {
				posList = append(posList[:p], posList[p+1:]...)
				break
			}
		}
		shuffledCards = append(shuffledCards, deck[cardPosition])
	}

	return shuffledCards
}
