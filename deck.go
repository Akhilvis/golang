package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"spades", "diamonds", "hearts", "club"}
	cardsValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, suit+"Of"+value)
		}
	}
	return cards
}
