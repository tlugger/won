package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
)

type Weather struct {
  Current Currently `json:"currently"`
}

type Currently struct {
  Summary string `json:"summary"`
  Temperature float64 `json:"temperature"`
}

func getRequest(url string) (*http.Response, error){
  resp, err := http.Get(url)
  return resp, err
}

func getWeather(body []byte) (*Weather, error){
  var w = new(Weather)
  err := json.Unmarshal(body, &w)

  return w, err
}

func main() {
  apikey := os.Getenv("WONAPIKEY")
  latLong := os.Getenv("WONLATLONG")

  s := []string{"https://api.darksky.net/forecast", apikey, latLong}
  url := strings.Join(s, "/")

  resp, err := getRequest(url)
  if err != nil{
    panic(err.Error())
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err.Error())
  }

  w, err := getWeather([]byte(body))
  if err != nil {
    panic(err.Error())
  }

  fmt.Printf("Current condition: %s\n", w.Current.Summary)
  fmt.Printf("Current temperature: %.2f°F\n", w.Current.Temperature)
}
