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

// TRANSLATE API TYPES
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

// Config (context) struct
type licences struct {
	TranslateApiKey string
	TwilioAccountSid string
	TwilioAuthToken string
	TwilioFromNumber string
}

// global config data
var gLicences *licences = &licences{}

// using: https://api.chucknorris.io/
func getJoke() (string, error) {
	// send request, get response
	res, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return "", err
	}

	log.Println("Received Response: " + res.Status)

	// read response body
	joke, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	log.Println("Received Response: " + string(joke))

	// extract the joke string from the response json object
	var jokeObj map [string]string

	// we can unmarshal this object as a map of string: string because the documentation tells us
	// that the response object is formatted in this way - otherwise, if you don't know the format
	// of the response object, you should not make assumptions as to how the data is stored.
	json.Unmarshal(joke, &jokeObj)
	if len(jokeObj["value"]) <= 0 {
		return "", errors.New("Failed to parse joke object")
	}

	// return the joke
	return jokeObj["value"], nil
}

// google translate request
func translateJoke(text string,	language string) (string, error) {
	queryParameters := &url.Values{
		"q": {text},
		"target": {language},
		"format": {"text"},
		"key": {globalConfig.TranslateApiKey},
	}

	res, err := http.Post("https://translation.googleapis.com/language/translate/v2?" +
		queryParameters.Encode(), "", nil)

	if err != nil {
		return "", err
	}

	// read response body
	translation, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	var translateObj TranslationResponse

	json.Unmarshal(translation, &translateObj)
	if len(translateObj.Data.Translations[0].TranslatedText) <= 0 {
		return "", errors.New("Failed to parse translation object")
	}

	return translateObj.Data.Translations[0].TranslatedText, nil
}

// Twilio request
func sendJoke(phoneNumber string, message string) error {
	twilioParams := &url.Values {
		"To": {phoneNumber},
		"From": {globalConfig.TwilioFromNumber},
		"Body": {message},
	}

	req, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/" +
		globalConfig.TwilioAccountSid + "/Messages.json", strings.NewReader(twilioParams.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(globalConfig.TwilioAccountSid, globalConfig.TwilioAuthToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	responseText, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	log.Println(string(responseText))

	return nil
}

func main() {

	// get program parameters
	phoneNumber := os.Args[1]
	jokeLanguage := os.Args[2]

	// set program global config variables
	globalConfig.TranslateApiKey = os.Args[3]
	globalConfig.TwilioAccountSid = os.Args[4]
	globalConfig.TwilioAuthToken = os.Args[5]
	globalConfig.TwilioFromNumber = os.Args[6]

	log.Println("To: " + phoneNumber + " From: " + globalConfig.TwilioFromNumber + " in: "+ jokeLanguage)

	var joke, translatedJoke string
	var err error

	// Get a joke
	joke, err = getJoke();
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Received Joke: " + joke)

	// Translate the joke to russian
	translatedJoke, err = translateJoke(joke, jokeLanguage)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Received Translation: " + translatedJoke)

	// Send the joke in an SMS
	err = sendJoke(phoneNumber, translatedJoke)
	if err != nil {
		log.Fatal(err.Error())
	}
}