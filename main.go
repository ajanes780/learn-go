package main

func main() {
	cards := newDeckFromFile("My Cards")

	//cards.print()
	cards.shuffle()
	cards.print()
}

/**
* Array -  is always fixed length
* Slice - is an array that can grow or shrink
 */
