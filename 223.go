package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://example.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
