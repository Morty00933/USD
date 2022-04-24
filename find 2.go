package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://example.com/")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	//	log.Printf(sb)
	sb2 := "This domain is for use in illustrative examples in documents. You may use this\n    domain in literature without prior coordination or asking for permission."
	if strings.Contains(sb, sb2) {
		fmt.Println(sb2)
	}
}
