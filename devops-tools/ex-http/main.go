package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	getExample1()
	getExample2()
	postExample1()

}

func getExample1() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("unable to create request")
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("unable to response")
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("unable to read body content")
	}

	fmt.Println(string(content))
}

func getExample2() {
	resp, err := http.Get("https://httpbin.org/get?search=params")
	if err != nil {
		log.Fatalln("unable to response")
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("unable to read body content")
	}

	fmt.Println(string(content))
}

func postExample1() {
	resp, err := http.Post("https://httpbin.org/post", "text/plan", strings.NewReader("this is a post request"))
	if err != nil {
		log.Fatalln("unable to response")
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("unable to read body content")
	}

	fmt.Println(string(content))
}
