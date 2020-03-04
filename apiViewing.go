// Giphy API
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func otherMain() {
	url := "https://giphy.p.rapidapi.com/v1/gifs/search?q=Get%20out%20of%20bed&api_key=dc6zaTOxFJmzC"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "giphy.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "72e63c5d87mshd02228ed7648d08p18efccjsn6164b5d330ec")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
