package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

/*
  Deep breath.... ok so, I want you to use at least 3 services from https://github.com/abhishekbanthia/Public-APIs

  I want you to use at least 2 HTTP request methods (GET and POST at minimum)

  I want you to use each of the following: Query Parameters, JSON request objects, x-www-form-urlencoded

  I want to see you writing golang structs for the json objects (see my example where I wrote the structs for the google translate API responses) ex:

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

  and, your program has to do something cool

  Bonus points for: trying XML / SOAP, etc. Trying APIs that use other HTTP methods (like: PUT, DELETE, etc),

  Also, they must use each of the following authentication mechanisms (1: no auth, 2: http basic auth, 3: authorization key (submitted via post or in query param))

  Bonus points if you can figure out how oauth works, and now to use it (not using a third party library, but instead using just HTTP, URL, and JSON)



  SUMMARY:
    3 API services
    2 HTTP methods GET & POST
    Use these 3:
      Query Parameters
      JSON request objects
      x-www-form-urlencoded
    Write JSON structs ALWAYYYSSSSS (see example above)
    3 types of authentication:
      no auth
      http basic auth
      authorization key via post or query param

  OPTIONS / IDEAS:
    Dictionary API
    NASA api?
    Rick and Morty Trivia
    Free texting APIs

    create a personal bot that listens for a request for information
    bot will respond to "wheres rick" to get a random location for Rick
    bot will respond to "nasa news" to get the latest news string from nasa
    bot will respond to "define" + user input <-- which will be a single word to define, the bot will respond with a definition if one exists.
    all 3 APIs connected to a message relay service (do i need a database for this?)
*/

/*   TEXTING API HERE

https://textbelt.com/

THIS IS JAVA TURN IT INTO GO

final NameValuePair[] data = {
    new BasicNameValuePair("phone", "5557727420"),
    new BasicNameValuePair("message", "Hello world"),
    new BasicNameValuePair("key", "textbelt")
};
HttpClient httpClient = HttpClients.createMinimal();
HttpPost httpPost = new HttpPost("https://textbelt.com/text");
httpPost.setEntity(new UrlEncodedFormEntity(Arrays.asList(data)));
HttpResponse httpResponse = httpClient.execute(httpPost);

String responseString = EntityUtils.toString(httpResponse.getEntity());
JSONObject response = new JSONObject(responseString);
*/

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
func wheresRick() (string, error) {

	resp, err := http.Get("https://rickandmortyapi.com/api/location/2")

	if err != nil {
		log.Fatal(err)
	}

	location, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var rickObj string

	// we can unmarshal this object as a map of string: string because the documentation tells us
	// that the response object is formatted in this way - otherwise, if you don't know the format
	// of the response object, you should not make assumptions as to how the data is stored.
	json.Unmarshal(location, &rickObj)
	if len(rickObj["name"]) <= 0 {
		return "", errors.New("Failed")
	}

	return rickObj["name"], nil

// RICK AND MORTY API END

func main() {

	globalConfig.textbeltKey = "xxxx"

	phone := os.Args[1]
	message, err := wheresRick()

	//send text message
	sendText(phone, message)

}
