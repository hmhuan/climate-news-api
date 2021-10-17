package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	var urls = []string{"https://theguardian.com/environment/climate-crisis"}

	resp, err := http.Get(urls[0])

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(body))

	htmlDoc, err := html.Parse(strings.NewReader(string(body)))

	fmt.Println(htmlDoc.Data)
}
