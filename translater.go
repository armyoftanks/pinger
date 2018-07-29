package main

import (
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"os"
	"fmt"
	"log"
	"context"
	"golang.org/x/text/language"
	"cloud.google.com/go/translate"

)

func main() {

	//GOOGLE TRANSLATE API STUFF HERE

	ctx := context.Background()

	gClient, errs := translate.NewClient(ctx)
	if errs != nil {
		log.Fatalf("Failed to create client: %v", errs)
	}

	text := os.Args[1]
	target, errs := language.Parse("ro")
	if errs != nil {
		log.Fatalf("Failed to parse target language: %v", errs)
	}

	// Translates the text into Romanian.
	translations, errs := gClient.Translate(ctx, []string{text}, target, nil)
	if errs != nil {
		log.Fatalf("Failed to translate text: %v", errs)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", translations[0].Text)

	//TWILIO API STUFF HERE

	accountSid := "x"
	authToken := "x"

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", os.Args[2])
	msgData.Set("From","x")
	msgData.Set("Body", translations[0].Text)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if (err == nil) {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status);
	}
}
