package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["black"] = "#000000"
	colors["white"] = "#ffffff"

	delete(colors, "white")

	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("The color '%v' is represented by '%v' hex \n", k, v)
	}
}
