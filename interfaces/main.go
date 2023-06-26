package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWritter struct {
}

func main() {
	// Example1()

	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Creating a custom Writes
	// io.Copy(os.Stdout, resp.Body)

	// watch out to create functions without responsabilities
	lw := logWritter{}
	io.Copy(lw, resp.Body)
}

func (l logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote this:", len(bs))

	return len(bs), nil
}
