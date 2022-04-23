package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.Get("https://example.com/")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

}
