package main

import "fmt"

type contactInfo struct {
	email      string
	postalCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jackson",
		contactInfo: contactInfo{
			email:      "a@a.com",
			postalCode: 123,
		},
	}
	jim.updateName("Tom")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
