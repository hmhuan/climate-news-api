package main

import (
	"log"
	"net/http"

	api "api"
)

func main() {
	var urls = []string{"https://theguardian.com/environment/climate-crisis"}

	resp, err := http.Get(urls[0])

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(body))

	// htmlDoc, err := html.Parse(strings.NewReader(string(body)))

	// fmt.Println(htmlDoc.Data)

	api.InitClimateAPI()

	api.HelloWold()

}
