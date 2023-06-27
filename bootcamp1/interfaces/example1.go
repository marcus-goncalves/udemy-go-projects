package main

import "fmt"

type bot interface {
	getGreeting() string
	// printGreeting()
}

type englishBot struct{}
type spanishBot struct{}

func Example1() {
	eb := englishBot{}
	printGreeting(eb)

	sb := spanishBot{}
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "Hello there"
}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
