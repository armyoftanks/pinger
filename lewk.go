package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
  "time"
  "os"
)

//connect twilio API
//connect to vision API
// user sends image to twilio
// image is analyzed by vision
// description of image is sent back to user

func analyzeImage(imageUrl string) {
  // For example, subscriptionKey = "0123456789abcdef0123456789ABCDEF"
    const subscriptionKey = "xxxx"

    // You must use the same location in your REST call as you used to get your
    // subscription keys. For example, if you got your subscription keys from
    // westus, replace "westcentralus" in the URL below with "westus".
    const uriBase =
        "https://westcentralus.api.cognitive.microsoft.com/vision/v2.0/analyze"

    const params = "?visualFeatures=Description&details=Landmarks&language=en"
    const uri = uriBase + params
    imageUrlEnc := "{\"url\":\"" + imageUrl + "\"}"

    reader := strings.NewReader(imageUrlEnc)

    // Create the Http client
    client := &http.Client{
        Timeout: time.Second * 2,
    }

    // Create the Post request, passing the image URL in the request body
    req, err := http.NewRequest("POST", uri, reader)
    if err != nil {
        panic(err)
    }

    // Add headers
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

    // Send the request and retrieve the response
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    // Read the response body.
    // Note, data is a byte array
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    // Parse the Json data
    var f interface{}
    json.Unmarshal(data, &f)

    // Format and display the Json result
    jsonFormatted, _ := json.MarshalIndent(f, "", "  ")
    fmt.Println(string(jsonFormatted))
}

func main()  {
  // picture is a image url string
  picture := os.Args[1]

  analyzeImage(picture)

}
