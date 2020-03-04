// Giphy API
package gif

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// gif represents the closest matching gif based on the user's input
type gif struct {
	Post []string `json:"post"`
}

func modifyInput(userInput string) string {
	input := userInput
	newInput := strings.ReplaceAll(input, " ", "%20")
	fmt.Println(input + " ==> What the user inputted,\n" + newInput + " ==> What the input is modified to.")
	return newInput
}

func searchGifs(userInput string) {
	// Modify userInput to be insertable into url variable

	url := "https://giphy.p.rapidapi.com/v1/gifs/search?q=" + modifyInput(userInput) + "&api_key=dc6zaTOxFJmzC"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "giphy.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "72e63c5d87mshd02228ed7648d08p18efccjsn6164b5d330ec")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	var theGif []string
	theGif = append(theGif, string(body))
	gif1 := gif{Post: theGif}
	fmt.Println(gif1)

	gif1JSON, _ := json.MarshalIndent(gif1, "", "    ")
	fmt.Println(string(gif1JSON))
}

func main() {
	searchGifs("Get out of bed")
}
