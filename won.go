package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// Weather unpacks the DarkSky data
type Weather struct {
	Current Currently `json:"currently"`
}

// Currently pulls summary and temperature from currently
type Currently struct {
	Summary     string  `json:"summary"`
	Temperature float64 `json:"temperature"`
}

func getRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	return resp, err
}

func getWeather(body []byte) (*Weather, error) {
	var w = new(Weather)
	err := json.Unmarshal(body, &w)

	return w, err
}

func getConditions() (string, float64) {
	apikey := os.Getenv("WONAPIKEY")
	latLong := os.Getenv("WONLATLONG")

	s := []string{"https://api.darksky.net/forecast", apikey, latLong}
	url := strings.Join(s, "/")

	resp, err := getRequest(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w, err := getWeather([]byte(body))
	if err != nil {
		log.Fatal(err)
	}

	return w.Current.Summary, w.Current.Temperature
}

func main() {
	current, temperature := getConditions()

	app := cli.NewApp()
	app.Name = "won"
	app.Usage = "Weather or not (here it comes)"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Current condition: %s\n", current)
		fmt.Printf("Current temperature: %.2f°F\n", temperature)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
