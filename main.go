package main

func main() {
	cards := newDeck()
	err := cards.saveToFile("My Cards")
	if err != nil {
		return
	}

}

/**
* Array -  is always fixed length
* Slice - is an array that can grow or shrink
 */
