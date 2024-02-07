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
	fmt.Printf("%+v", jim)
}
