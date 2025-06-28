package main

import "fmt"

func war(pCard1, pCard2 player) (player, error) {
	for i := 0; i < 10; i++ { // temporary fixed loop to not iterate over the whole game
		if len(pCard1.hand) == 0 {
			if len(pCard1.pile) == 0 {
				fmt.Println("PLAYER 1 IS THE WINNER")
				return pCard1, nil
			} else {
				pCard1.hand = shuffle(pCard1.pile)
				pCard1.pile = pCard1.pile[:0]
			}
		}
		if len(pCard2.hand) == 0 {
			if len(pCard2.pile) == 0 {
				println("PLAYER 2 IS THE WINNER")
				return pCard2, nil
			} else {
				pCard2.hand = shuffle(pCard1.pile)
				pCard2.pile = pCard2.hand[:0]
			}
		}
		handWinner := ""
		if pCard1.hand[0].value > pCard2.hand[0].value {
			pCard2.pile = append(pCard2.pile, pCard1.hand[0])
			pCard2.pile = append(pCard2.pile, pCard2.hand[0])
			handWinner = "player 1"
		} else {
			pCard1.pile = append(pCard1.pile, pCard1.hand[0])
			pCard1.pile = append(pCard1.pile, pCard2.hand[0])
			handWinner = "player 2"
		}
		fmt.Printf("GAME STATUS:\n player1 card: %v player2 card: %v Roundwinner: %s\n ", pCard1.hand[0].name, pCard2.hand[0].name, handWinner)
		pCard1.hand = pCard1.hand[1:]
		pCard2.hand = pCard2.hand[1:]
		fmt.Printf("player 1 status, deck: %v, pile %v\n ", len(pCard1.hand), len(pCard1.pile))
		fmt.Printf("player 2 status, deck: %v, pile %v\n ", len(pCard2.hand), len(pCard2.pile))
	}
	// temporary logic to not debug entire gameplay length
	player1CardsLeft := len(pCard1.hand) + len(pCard1.pile)
	player2CardsLeft := len(pCard2.hand) + len(pCard2.pile)
	if player1CardsLeft > player2CardsLeft {
		return pCard2, nil
	} else {
		return pCard1, nil
	}
}
