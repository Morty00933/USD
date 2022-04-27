package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://cbr.ru/key-indicators/")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	//	log.Printf(sb)
	//	sb2 := "This domain is for use in illustrative examples in documents. You may use this\n    domain in literature without prior coordination or asking for permission."
	//	if strings.Contains(sb, "12344532") {
	//		fmt.Println(sb2)
	//	}
	//	matched, _ := regexp.MatchString(`[1-9]+[0-9]+\,+[0-9]{4}`, sb)
	//	fmt.Println(matched)
	re := regexp.MustCompile(`[1-9]+[0-9]+\,+[0-9]{4}`)
	fmt.Printf("%q\n", re.FindString(sb)) // "food"
}
