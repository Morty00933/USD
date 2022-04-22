package main

import (
	"bufio"
	"fmt"
	"net/http"
	_ "strings"
)

func main() {

	resp, err := http.Get("https://example.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for i := 20; scanner.Scan() && i < 50; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
