package main

import (
	"net/http"
	"net/url"
	"log"
	"fmt"
	"encoding/json"
)

/*
OBJECTIVE: TEXT YOUR FRIENDS IN A DIFF LANGUAGE
1. SEND MESSAGE TO GOOGLE TRANSLATE
2. RECEIVE TRANSLATION AND STORE INTO MESSAGE
3. SEND MESSAGE TO FRIENDS
*/


func main () {
	// ME ATTEMPTING SOMETHING post + http + query parameters
	// YOU NEED TO CREATE AN HTTP REQUEST WITH URL QUERY STRINGS OF THE TRANSLATE TYPES SPECIFIED I THINK
	// GOOGLE TRANSLATE
	client := &http.Client{}
	gparameters := url.Values{}
	gparameters.Add("q", "hello world")
	gparameters.Add("target", "ru")
	gparameters.Add("source", "en")
	msgDataReader,_ := http.NewRequest("POST", "https://translation.googleapis.com/language/translate/v2" + gparameters.Encode(), nil)
	msgDataReader.Header.Add("content-type", "application/json")

	resp, err := client.Do(msgDataReader)
	if err != nil{
		log.Fatalf("Failed to translate text: %v", err)
	}
	fmt.Println(resp)

	if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err2 := decoder.Decode(&data)
		if (err2 == nil) {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status);
	}

	/* TWILIO
	accountSid := "xxxx"
	authToken := "xxxx"
	urlStr2 := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	msgData2 := url.Values{}
	msgData2.Set("To", os.Args[1])
	msgData2.Set("From","xxxx")
	msgData2.Set("Body", os.Args[2])  // <---  I WANT TO REPLACE THIS WITH THE TRANSLATED TEXT FROM GOOGLE TRANSLATE RESPONSE ABOVE
	msgDataReader2 := *strings.NewReader(msgData2.Encode())
	client2 := &http.Client{}
	req2, _ := http.NewRequest("POST", urlStr2, &msgDataReader2)
	req2.SetBasicAuth(accountSid, authToken)
	req2.Header.Add("Accept", "application/json")
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp2, _ := client2.Do(req2)
	if (resp2.StatusCode >= 200 && resp2.StatusCode < 300) {
		var data2 map[string]interface{}
		decoder2 := json.NewDecoder(resp2.Body)
		err2 := decoder2.Decode(&data2)
		if (err2 == nil) {
			fmt.Println(data2["sid"])
		}
	} else {
		fmt.Println(resp2.Status);
	}
*/

}
