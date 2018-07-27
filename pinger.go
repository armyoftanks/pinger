package main

import (
	"net/url"
	"os"
	"strings"
	"net/http"
	"encoding/json"
	"fmt"
)
type detectedSourceLanguage struct {
	target string
}

type translatedText struct {
	q string
}



func main() {

	//TWILIO API STUFF HERE

	accountSid := "xxxx"
	authToken := "xxxx"

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", os.Args[1])
	msgData.Set("From","xxxx")
	msgData.Set("Body", os.Args[2])
	msgData.Set("Target", os.Args[3])
	msgData.Set("Model", "nmt")
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


	//GOOGLE TRANSLATE API STUFF HERE
	translation := {
		detectedSourceLanguage{
			target: msgData.Target(),
		}
	}

	req, _ := http.NewRequest("POST", "https://translation.googleapis.com/language/translate/v2", translation.traslatedText())

}