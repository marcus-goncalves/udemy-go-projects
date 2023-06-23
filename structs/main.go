package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	alex := Person{"alex", "andersen"}

	jean := Person{
		firstName: "Jean",
		lastName:  "Billy",
	}

	var loris Person
	loris.firstName = "Loris"
	loris.lastName = "Lane"

	fmt.Println(alex, jean, loris)

}
