package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

}

/**
* Array -  is always fixed length
* Slice - is an array that can grow or shrink
 */
