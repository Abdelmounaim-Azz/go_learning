package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	john := person{firstName: "John", lastName: "Doe"}
	fmt.Printf("%+v", john)
	john.firstName = "John2"
	john.lastName = "Doe2"
	fmt.Printf("%+v", john)
}
