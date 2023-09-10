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
	Length int    `json:"length"`
}

type randomUser struct {
	Results []userResult
}

type userResult struct {
	Gender  string
	Name    userName
	Email   string
	Picture userPicture
}

type userName struct {
	First string
	Last  string
}

type userPicture struct {
	Large  string
	Medium string
}

func getCatFact() {

	url := "https://catfact.ninja/fact"
	var r catfact
	err := getJson(url, &r)
	if err != nil {
		fmt.Printf("Error getting catfact %s", err)
	} else {
		fmt.Println("A super intersting catfact", r.Fact)

	}
}

func getRandomUser() {
	url := "https://randomuser.me/api/"
	var user randomUser
	err := getJson(url, &user)
	if err != nil {
		fmt.Printf("Error getting random user %s", err)
	} else {
		fmt.Printf("User info: \nGender: %s \nFirst Name: %s \nLast Name: %s \nPicture: %s\n",
			user.Results[0].Gender,
			user.Results[0].Name.First,
			user.Results[0].Name.Last,
			user.Results[0].Picture.Medium)
	}
}

func getJson(url string, target interface{}) error {

	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	//var r catfact

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&target); err != nil {
		return err
	}
	return nil

}

func main() {

	client = &http.Client{Timeout: 10 * time.Second}

	getCatFact()
	getRandomUser()

	catfact2 := catfact{
		Fact:   "Random fact",
		Length: 12,
	}

	jsonStr, err := json.Marshal(catfact2)
	if err != nil {
		fmt.Println("Unable to Marshal", err)
	} else {
		fmt.Printf("The marshaled string is: %s\n", string(jsonStr))
	}

}
