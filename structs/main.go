package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	fmt.Println(alex)
	alex.firstName = "Alex"
	alex.lastName = "Andersen"
	fmt.Printf("%+v", alex)
}
