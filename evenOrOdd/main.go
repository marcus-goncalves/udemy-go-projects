package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ {
		checkNumber(i)
	}
}

func checkNumber(i int) {
	if i%2 == 0 {
		fmt.Println(i, " - even")
	} else {
		fmt.Println(i, " - odd")
	}
}
