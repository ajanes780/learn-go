package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("======")
	remainingCards.print()
}

/**
* Array -  is always fixed length
* Slice - is an array that can grow or shrink
 */
