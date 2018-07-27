package main

import (
	"fmt"
	//"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	//"log"
	"encoding/json"
	"os"
)

func main() {

	accountSid := "xxxx"
	authToken := "xxxx"

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	ticker := time.NewTicker(time.Hour)

	func chuck := {
		if (ticker % 10)
		http.NewRequest("GET", https://api.chucknorris.io/jokes/random)
	}

	msgData := url.Values{}
	msgData.Set("To", os.Args[1])
	msgData.Set("From","xxxx")
	msgData.Set("Body", chuck)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//i get this part
	resp, _ := client.Do(req)
	if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
	  var data map[string]interface{}
		//pretty much start losing me here
		decoder := json.NewDecoder(resp.Body)
		//what is data pointing to?
	  err := decoder.Decode(&data)
	  if (err == nil) {
	    fmt.Println(data["sid"])
	  }
	} else {
	  fmt.Println(resp.Status);
	}

	/* fmt.Println(os.Args[2])
	s, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))  */

}
