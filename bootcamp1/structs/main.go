package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	alex := Person{"alex", "andersen", ContactInfo{"test@test.io", 30882570}}

	jean := Person{
		firstName: "Jean",
		lastName:  "Billy",
		contact: ContactInfo{
			email:   "test.test.io",
			zipCode: 30882570,
		},
	}

	var loris Person
	loris.firstName = "Loris"
	loris.lastName = "Lane"
	loris.contact.email = "test@test.io"
	loris.contact.zipCode = 30882570

	fmt.Println(alex, jean, loris)
	alex.print()
	jean.changeName("Jeany")
	jean.print()
}

func (p Person) print() {
	fmt.Println(p.firstName)
}

func (p *Person) changeName(newName string) {
	p.firstName = newName
}
