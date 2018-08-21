package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Config (context) struct
type appConfig struct {
	textbeltKey string
}

// global config data
var globalConfig *appConfig = &appConfig{}

// TEXT BELT API HERE
func sendText(phone string, message string) (string, error) {

	resp, err := http.PostForm("https://textbelt.com/text", url.Values{"phone": {phone}, "message": {message}, "key": {globalConfig.textbeltKey}})
	if err != nil {
		return "", err
	}

	textbelt, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	log.Println(string(textbelt))

	return string(textbelt), err
}

// TEXT BELT API END

// RICK AND MORTY API HERE

type locationInfo struct {
	id        int            `json:"id"`
	name      string         `json:"name"`
	types     string         `json:"type"`
	dimension string         `json:"dimention"`
	residents []residentList `json:"residents"`
	url       string         `json:"url"`
	created   string         `json:"created"`
}

type residentList struct {
	residents string
}

func wheresRick() (string, error) {

	url := "https://rickandmortyapi.com/api/location/2"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var l locationInfo

	json.Unmarshal(body, &l)
	fmt.Printf("Results: %v\n", l)

	return string(l.name), err
}

// RICK AND MORTY API END

func main() {

	globalConfig.textbeltKey = "textbelt"

	phone := os.Args[1]
	message, _ := wheresRick()
	fmt.Println(message)
	//send text message
	sendText(phone, message)

}
