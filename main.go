package main

// import "fmt"

func main() {

	// cards := newDeck()
	// cards.writeFile("my_cards")

	// fmt.Println(cards.toString())

	// hand, remaingCard := deal(cards, 5)

	// hand.print()
	// // fmt.Print("##############################################")
	// remaingCard.print()

	cards := newDeckFromFile("my_cards")

	cards.print()
}

func newCard() string {
	return "five of dimonds"
}
