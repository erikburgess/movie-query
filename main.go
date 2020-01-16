package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {

	const apiUrl = "https://www.omdbapi.com/?apikey=[apiKey]&t=[movieTitle]"

	movieName := flag.String("name", "", "Name of the movie to query")

	flag.Parse()

	apiKey, found := os.LookupEnv("OMDB_API_KEY")

	if !found {
		log.Fatalf("Missing OMDB_API_KEY env variable")
	}

	completeUrl := strings.Replace(apiUrl, "[apiKey]", apiKey, 1)

	escapedMovieName := url.QueryEscape(*movieName)

	completeUrl = strings.Replace(completeUrl, "[movieTitle]", escapedMovieName, 1)

	resp, err := http.Get(completeUrl)

	if err != nil {
		log.Fatalf("Failed to call api %v", err.Error())
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Failed to read response body %v", err.Error())
	}

	var parsedResponseBody map[string]interface{}

	err = json.Unmarshal(responseBody, &parsedResponseBody)

	if err != nil {
		log.Fatalf("Failed to parse response body %v", err.Error())
	}

	//log.Printf("Response %v", parsedResponseBody)

	responseResult, err := strconv.ParseBool(parsedResponseBody["Response"].(string))

	if err != nil {
		log.Fatalf("Invalid response %v", err.Error())
	}

	if !responseResult {
		log.Fatalf("%v", parsedResponseBody["Error"])
	}

	ratings := parsedResponseBody["Ratings"].([]interface{})

	ratingFound := false
	for _, rating := range ratings {
		if rating.(map[string]interface{})["Source"] == "Rotten Tomatoes" {
			ratingFound = true
			fmt.Printf("%v", rating.(map[string]interface{})["Value"])
			break
		}
	}

	if !ratingFound {
		log.Fatalf("Could not find Rotten Tomato rating for movie")
	}
}
