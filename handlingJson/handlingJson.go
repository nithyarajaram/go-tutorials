package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type catfact struct {
	Fact   string `json:"fact"`
	Length int
}

func getCatFact() (string, error) {

	url := "https://catfact.ninja/fact"
	fact, err := getJson(url)
	catfact := fmt.Sprintf(fact)
	return catfact, err

}

func getJson(url string) (string, error) {

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var r catfact

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", err
	}
	catfact := fmt.Sprintf(r.Fact)
	return catfact, nil

}

func main() {

	client = &http.Client{Timeout: 10 * time.Second}

	catfact, err := (getCatFact())
	if err != nil {
		fmt.Printf("Could not get catfact %d", err)
	}
	fmt.Println(catfact)

}
