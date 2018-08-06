package main

import (
	"log"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

type TranslateTextResponseTranslation struct {
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
	Model string `json:"model"`
	TranslatedText string `json:"translatedText"`
}

type TranslateTextResponseList struct {
	Translations []TranslateTextResponseTranslation `json:"translations"`
}

type TranslationResponse struct {
	Data TranslateTextResponseList `json:"data"`
}

type gLicences struct {
	TranslateApiKey string
	TwilioAccountSid string
	TwilioAuthToken string
	TwilioFromNumber string
}

// secuirty data
var gLicences *licences = &licences{}

func getJoke() (string, error) {
	res, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nill {
		return "", err
	}

	log.Println("Received Response: " + res.Status)

	var jokeObj map [string]string

	json.Unmarshal(joke, &jokeObj)

	if len(jokeObj["value"]) <= 0 {
		return "", errors.New("Failed to parse joke object")
	}

	return jokeObj["value"], nil
}

func main() {

	gLicences.TranslateApiKey = os.Args[3]
	gLicences.TwilioAccountSid = os.Args[4]
	gLicences.TwilioAuthToken = os.Args[5]
	gLicences.TwilioFromNumber = os.Args[6]

	joke, err := getJoke();
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Received Joke: " + joke)
}