package main

import (
	"bytes"
	"encoding/base64"
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

	// Create a base64 encoded user and password
	buffer := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	enc.Write([]byte("user:passw0rd"))
	enc.Close()
	creds, err := buffer.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		log.Fatalln("unable to encode credentials")
	}

	// Add header
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", creds))

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
