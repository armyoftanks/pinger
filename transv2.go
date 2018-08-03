package main

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"encoding/json"
)

/*
OBJECTIVE: TEXT YOUR FRIENDS A RANDOM CHUCK NORRIS JOKE IN A DIFF LANGUAGE
1. GET RANDOM CHUCK NORRIS JOKE
2. TRANSLATE JOKE TO WHATEVER LANGUAGE YOU CHOOSE
3. SEND TRANSLATED JOKE AS MESSAGE TO FRIENDS
*/

type joker string

func main () {
	//CHUCK NORRIS RANDOM JOKE API
	res,_ := http.Get("https://api.chucknorris.io/jokes/random")
	if (res.StatusCode >= 200 && res.StatusCode < 300) {
		var data map[string]interface{}
		decoder := json.NewDecoder(res.Body)
		err := decoder.Decode(&data)
		if (err == nil) {
			fmt.Println(data["value"])
		}
	} else {
		fmt.Println(res.Status);
	}

	joke,_ := ioutil.ReadAll(res.Body)


	// GOOGLE TRANSLATE API TRANSLATING CHUCK NORRIS JOKE
	gparameters := url.Values{}
	gparameters.Add("q", string(joke))
	gparameters.Add("target", os.Args[1])
	gparameters.Add("source", "en")
	gparameters.Add("key", "xxxxx")
	resp,_:= http.Post("https://translation.googleapis.com/language/translate/v2?" + gparameters.Encode(), "application/json", nil)

	responseText, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(responseText))


	//TWILIO
	accountSid := "xxxxx"
	authToken := "xxxxx"
	urlStr2 := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	msgData2 := url.Values{}
	msgData2.Set("To", os.Args[2])
	msgData2.Set("From",	"xxxxx")
	msgData2.Set("Body", string(responseText))
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

}
