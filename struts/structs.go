package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		age:       32,
	}
	fmt.Printf("%+v", alex)
	alex.firstName = "Aaron"
	fmt.Printf("%+v", alex)
}
