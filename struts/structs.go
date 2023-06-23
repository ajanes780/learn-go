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
	jim.updateName("Aaron")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
