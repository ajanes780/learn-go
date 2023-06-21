package main

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")

	cards.print()
	//fmt.Println(cards)
}

/**
* Array -  is always fixed length
* Slice - is an array that can grow or shrink
 */
