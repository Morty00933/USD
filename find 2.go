package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	resp, err := http.Get("https://www.sberometer.ru/?c&&&&")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	re := regexp.MustCompile(`[1-9]+[0-9]+\.+[0-9]{4}`)
	fmt.Printf("%q\n", re.FindString(sb))
	RU, err := strconv.ParseFloat(re.FindString(sb), 32)
	var EUR float64
	var ru int = int(RU)
	eur, err := fmt.Scanln(EUR)
	if err != nil {
		return
	}
	sum := ru * eur
	fmt.Println(sum)
}
