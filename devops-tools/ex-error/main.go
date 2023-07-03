package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var router = NewRouter()

func init() {
	router.Register(200, func(resp *http.Response) {
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("unable to get content")
		}

		fmt.Println(string(content))
	})

	router.Register(404, func(r *http.Response) {
		log.Fatalln("Not Found:", r.Request.URL.String())
	})
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln("unable to read response")
	}

	router.Process(resp)
}
