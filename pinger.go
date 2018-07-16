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

	accountSid := "XXXX"
	authToken := "XXXX"

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", os.Args[1])
	msgData.Set("From","+XXXX")
	msgData.Set("Body", os.Args[2])
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

}
