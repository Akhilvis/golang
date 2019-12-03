package main

import "fmt"

// import "fmt"

func main() {

	cards := newDeck()

	fmt.Println(cards)

	hand, remaingCard := deal(cards, 5)

	hand.print()
	fmt.Print("##############################################")
	remaingCard.print()

}

func newCard() string {
	return "five of dimonds"
}
