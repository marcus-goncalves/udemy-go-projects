package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	if err != nil {
		log.Fatalln("unable to create request")
	}

	// easier way
	req.SetBasicAuth("user", "passw0rd")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("unable to get response")
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("unable to read response")
	}

	fmt.Println(string(content))
	fmt.Println(resp.StatusCode)

}
