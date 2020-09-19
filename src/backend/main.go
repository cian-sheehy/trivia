package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Trivia results from API
type Trivia struct {
	Results []struct {
		Question         string   `json:"question"`
		CorrectAnswer    string   `json:"correct_answer"`
		IncorrectAnswers []string `json:"incorrect_answers"`
	} `json:"results"`
}

// Field for slack
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Home endpoint")
	resp, err := http.Get("https://opentdb.com/api.php?amount=20&category=9&difficulty=medium&type=multiple")

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

	var t Trivia

	jsonErr := json.Unmarshal(body, &t)

	if jsonErr != nil {
		log.Fatalf("unable to parse value: %q, error: %s", string(body), jsonErr.Error())
	}

	log.Println(t)
	json.NewEncoder(w).Encode(t)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/question", question)
	http.HandleFunc("/answer", answer)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}
