package main

import "fmt"

func startWar(pCard1, pCard2 player, pile []cards) (player, player) {
	if len(pCard1.hand) <= 3 {
		pCard1.hand = append(pCard1.hand, shuffle(pCard1.pile)...)
		pCard1.pile = pCard1.pile[:0]
		if len(pCard1.hand) <= 3 {
			pCard1.hand = pCard1.hand[:0]
			return pCard1, pCard2
		}
	}
	if len(pCard2.hand) <= 3 {
		pCard2.hand = append(pCard2.hand, shuffle(pCard2.pile)...)
		pCard2.pile = pCard2.pile[:0]
		if len(pCard2.hand) <= 3 {
			pCard2.hand = pCard2.hand[:0]
			return pCard1, pCard2
		}
	}
	pile = append(pile, pCard1.hand[:3]...)
	pile = append(pile, pCard2.hand[:3]...)
	// updates a pile where the pot is created
	//
	if pCard1.hand[3].value < pCard2.hand[3].value {
		pCard1.hand = pCard1.hand[3:] // removes cards from hand
		pCard2.hand = pCard2.hand[3:]
		pCard1.hand = append(pCard1.hand, pile...) // adds card from pile to the loser
		println("PLAYER 2 IS THE WINNER")
		return pCard1, pCard2
	}
	if pCard1.hand[3].value > pCard2.hand[3].value {
		pCard1.hand = pCard1.hand[3:] // removes cards from hand
		pCard2.hand = pCard2.hand[3:]
		pCard2.hand = append(pCard2.hand, pile...) // adds card from pile to the loser
		println("PLAYER 1 IS THE WINNER")
		return pCard1, pCard2
	}
	if pCard1.hand[3].value == pCard2.hand[3].value {
		pCard1.hand = pCard1.hand[3:] // removes cards from hand
		pCard2.hand = pCard2.hand[3:]
		startWar(pCard1, pCard2, pile)
	}
	return pCard1, pCard2
}

func war(pCard1, pCard2 player) (player, error) {
	for { // temporary fixed loop to not iterate over the whole game
		//checks if players still have cards
		if len(pCard1.hand) == 0 {
			if len(pCard1.pile) == 0 {
				fmt.Println("PLAYER 1 IS THE WINNER")
				return pCard1, nil
			} else {
				pCard1.hand = append(pCard1.hand, shuffle(pCard1.pile)...)
				pCard1.pile = pCard1.pile[:0]
			}
		}
		if len(pCard2.hand) == 0 {
			if len(pCard2.pile) == 0 {
				println("PLAYER 2 IS THE WINNER")
				return pCard2, nil
			} else {
				pCard2.hand = append(pCard2.hand, shuffle(pCard2.pile)...)
				pCard2.pile = pCard2.hand[:0]
			}
		}
		// keep track of who wins and what card are played
		handWinner := ""
		player1Card := pCard1.hand[0].name
		player2Card := pCard2.hand[0].name
		// does it start a war, if a player enters a was with less than 3 cards he wins
		if pCard1.hand[0].value == pCard2.hand[0].value {
			fmt.Println("\nSTARTING A WAR ")
			pile := []cards{}
			pCard1, pCard2 = startWar(pCard1, pCard2, pile)
		} else if pCard1.hand[0].value > pCard2.hand[0].value {
			pCard2.pile = append(pCard2.pile, pCard1.hand[0])
			pCard2.pile = append(pCard2.pile, pCard2.hand[0])
			handWinner = "player 1"
			pCard1.hand = pCard1.hand[1:]
			pCard2.hand = pCard2.hand[1:]
		} else {
			pCard1.pile = append(pCard1.pile, pCard1.hand[0])
			pCard1.pile = append(pCard1.pile, pCard2.hand[0])
			pCard1.hand = pCard1.hand[1:]
			pCard2.hand = pCard2.hand[1:]
			handWinner = "player 2"
		}
		fmt.Printf("GAME STATUS:\n player1 card: %v player2 card: %v Roundwinner: %s\n ", player1Card, player2Card, handWinner)

		fmt.Printf("player 1 status, deck: %v, pile %v\n ", len(pCard1.hand), len(pCard1.pile))
		fmt.Printf("player 2 status, deck: %v, pile %v\n ", len(pCard2.hand), len(pCard2.pile))
	}
	// temporary logic to not debug entire gameplay length
	/*player1CardsLeft := len(pCard1.hand) + len(pCard1.pile)
	player2CardsLeft := len(pCard2.hand) + len(pCard2.pile)
	if player1CardsLeft > player2CardsLeft {
		return pCard2, nil
	} else {
		return pCard1, nil
	}*/
}
