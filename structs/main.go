package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "john.doe@doe.com",
			zipCode: 12345,
		},
	}
	joinP := &john
	joinP.updateFirstName("Johnny")
	john.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)

}
