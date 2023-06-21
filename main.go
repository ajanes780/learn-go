package main

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
	//fmt.Println(cards)
}

func newCard() string {
	return "Five of diamonds"
}

/**
* Array -  is always fixed length
* Slice - is an array that can grow or shrink
 */
