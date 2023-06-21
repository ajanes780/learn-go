package main

import "fmt"

// create a new type of deck
// Which is a slice of strings

type deck []string

// receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
