package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"trivia/backend/constants"
	"trivia/backend/types"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Home endpoint")
	var amount = "50"
	var difficulty = "medium"
	resp, err := http.Get(constants.BaseURL + "?amount=" + amount + "&category=9&difficulty=" + difficulty + "&type=multiple")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("HTTP: %s\n", resp.Status)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var t types.Trivia

	jsonErr := json.Unmarshal(body, &t)

	if jsonErr != nil {
		log.Fatalf("unable to parse value: %q, error: %s", string(body), jsonErr.Error())
	}
	fmt.Println("Sending back data...")

	for i := 0; i < len(t.Results); i++ {
		t.Results[i].IncorrectAnswers = append(t.Results[i].IncorrectAnswers, t.Results[i].CorrectAnswer)
	}

	json.NewEncoder(w).Encode(t)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
