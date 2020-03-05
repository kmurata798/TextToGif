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
	// Modify userInput to be insertable into url variable
	input := userInput
	// replace all replaces target char/string with desired char/string
	newInput := strings.ReplaceAll(input, " ", "%20")
	fmt.Println(input + " ==> What the user inputted,\n" + newInput + " ==> What the input is modified to.")
	return newInput
}
func between(value string, a string, b string) string {
	// Get substring between two strings.
	startingPoint := strings.Index(value, a)
	// strings.Index() give the index where the target string, a, exists in the given string, value
	if startingPoint == -1 {
		return ""
	}
	endPoint := strings.Index(value, b)
	// same as startingPoint variable, but gives the endpoint string index
	if endPoint == -1 {
		return ""
	}
	subStringStartingPoint := startingPoint + len(a)
	/* This variable is assigned the index where the starting point string begins, then adding the length of the starting point string
	which is the the starting point of the target string we are trying to grab.
	*/
	if subStringStartingPoint >= endPoint {
		return ""
	}
	// from the given string, return the substring in between the startingPointAdjusted
	return value[subStringStartingPoint:endPoint]
}

func searchGifs(userInput string) {
	// Implementing GIPHY API
	// ==BEGINNING OF GIPHY API==
	// Modifying requested GIPHY API url to allow user input to dynamically be inserted into the url.
	url := "https://giphy.p.rapidapi.com/v1/gifs/search?q=" + modifyInput(userInput) + "&api_key=dc6zaTOxFJmzC"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "giphy.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "72e63c5d87mshd02228ed7648d08p18efccjsn6164b5d330ec")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	// ==END OF GIPHY API==

	cleanedGifUrl := strings.ReplaceAll(string(body), "\\", "")
	// fmt.Println(replacedGif1)
	cleanedGifUrl2 := strings.ReplaceAll(cleanedGifUrl, "\"", "")
	// fmt.Println(replacedGif2)
	// findUrl :=
	fmt.Println(between(cleanedGifUrl2, "url:", ",slug:"))
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
