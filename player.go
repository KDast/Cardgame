package main

type player struct {
	hand []cards
	name string

	pile []cards
	//card to play are in hand, cards player picks up goes into pile
}
