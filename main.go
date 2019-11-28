package main

import "fmt"

// import "fmt"

func main() {

	cards := newDeck()

	fmt.Println(cards)

}

func newCard() string {
	return "five of dimonds"
}
