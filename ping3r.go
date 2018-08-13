package main

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
*/

func main() {

}
