package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
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
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Dimension string   `json:"dimention"`
	Residents []string `json:"residents"`
	Url       string   `json:"url"`
	Created   string   `json:"created"`
}

func wheresRick() (string, error) {
	rand.Seed(time.Now().Unix())
	locationID := (rand.Int() % 76) + 1
	url := fmt.Sprintf("https://rickandmortyapi.com/api/location/%d", locationID)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}

	var l locationInfo

	json.Unmarshal(body, &l)
	return "HackerQween says Ricks current location is: " + string(l.Name), nil
}

// RICK AND MORTY API END

func main() {

	globalConfig.textbeltKey = "xxxx"

	phone := os.Args[1]
	message, _ := wheresRick()
	fmt.Println(message)
	sendText(phone, message)

}
