package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
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
		contact: contactInfo{
			email:   "jimparty@gmail.com",
			phone:   7807201385,
			zipcode: 90210,
		},
	}
	fmt.Printf("%+v", jim)
}
