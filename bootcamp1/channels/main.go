package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go getLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			getLink(link, c)
		}(l)
	}
}

func getLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
