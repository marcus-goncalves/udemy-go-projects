package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GetResponse struct {
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func main() {
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

	txtContent := &GetResponse{}
	json.Unmarshal(content, txtContent)
	fmt.Println(txtContent.ToString())
}

func (r *GetResponse) ToString() string {
	out := fmt.Sprintf("GET RESPONSE:\nOrigin IP: %s\nRequest URL: %s", r.Origin, r.URL)
	for k, v := range r.Headers {
		out += fmt.Sprintf("Header %s: %s\n", k, v)
	}

	return out
}
