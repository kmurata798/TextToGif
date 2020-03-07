// Giphy API
package gif

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// gif represents the closest matching gif based on the user's input
type gif struct {
	Post []string `json:"post"`
}

func ModifyInput(userInput string) string {
	// Modify userInput to be insertable into url variable
	input := userInput
	// replace all replaces target char/string with desired char/string
	newInput := strings.ReplaceAll(input, " ", "%20")
	fmt.Println(input + " ==> User's input,\n" + newInput + " ==> Modified input.")
	return newInput
}
func Between(given string, first string, second string) string {
	// Get substring between two strings.
	startingPoint := strings.Index(given, first)
	// strings.Index() give the index where the target string, a, exists in the given string, value
	if startingPoint == -1 {
		return ""
	}
	endPoint := strings.Index(given, second)
	// same as startingPoint variable, but gives the endpoint string index
	if endPoint == -1 {
		return ""
	}
	subStringStartingPoint := startingPoint + len(first)
	/* This variable is assigned the index where the starting point string begins, then adding the length of the 
	starting point string which is the the starting point of the target string we are trying to grab.
	*/
	if subStringStartingPoint >= endPoint {
		return ""
	}
	// from the given string, return the substring in between the startingPointAdjusted
	fmt.Println("First bounds: " + first + "\nSecond bounds: " + second)
	fmt.Println("Target substring ==> " + given[subStringStartingPoint:endPoint] + "\n")
	return given[subStringStartingPoint:endPoint]
}

func SearchGifs(userInput string) string {
	// Implementing GIPHY API
	// ==BEGINNING OF GIPHY API==
	// Modifying requested GIPHY API url to allow user input to dynamically be inserted into the url.
	url := "https://giphy.p.rapidapi.com/v1/gifs/search?q=" + ModifyInput(userInput) + "&api_key=dc6zaTOxFJmzC"

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

	fmt.Println(Between(cleanedGifUrl2, "url:", ",slug:"))
	gifURL := Between(cleanedGifUrl2, "url:", ",slug:")
	return gifURL
}
