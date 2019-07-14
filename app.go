package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type peopleList struct {
	PersonList []person `json:"people"`
	Number     int      `json:"number"`
	Message    string   `json:"message"`
}

type person struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

func main() {
	url := "http://api.open-notify.org/astros.json"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var people peopleList

	jsonErr := json.Unmarshal(body, &people)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for i := 0; i < len(people.PersonList); i++ {
		fmt.Printf("%s is aboard the %s\n", people.PersonList[i].Name, people.PersonList[i].Craft)
	}
	os.Exit(0)
}
