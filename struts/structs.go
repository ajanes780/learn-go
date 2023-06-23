package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

type contactInfo struct {
	email   string
	phone   int
	zipcode int
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		age:       32,
		contactInfo: contactInfo{
			email:   "jimparty@gmail.com",
			phone:   7807201385,
			zipcode: 90210,
		},
	}
	// creates a pointer to the value
	jimPointer := &jim
	jimPointer.updateName("Aaron")
	jimPointer.print()
}

//// this is passed by value - copy from memory Go is a pass by value language
//func (p person) print() {
//	fmt.Printf("%+v", p)
//}

//func (p person) updateName(newFirstName string) {
//	p.firstName = newFirstName
//}

// Uses a pointer to now it is passed by reference
func (p *person) print() {
	fmt.Printf("%+v", p)
}

// Uses a pointer to now it is passed by reference
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
