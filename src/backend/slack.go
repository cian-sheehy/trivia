package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Action for slack
type Action struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Text  string `json:"text"`
	Style string `json:"style"`
}

// Attachment for slack
type Attachment struct {
	AttachmentType string `json:"attachment_type"`
	Color          string `json:"color"`
	// PreText string `json:"pretext"`
	// AuthorName string  `json:"author_name"`
	// AuthorLink string  `json:"author_link"`
	// AuthorIcon string  `json:"author_icon"`
	// Title string `json:"title"`
	// TitleLink string  `json:"title_link"`
	Text *string `json:"text"`
	// ImageUrl string  `json:"image_url"`
	// Fields []Field `json:"fields"`
	// Footer     string  `json:"footer"`
	// FooterIcon string  `json:"footer_icon"`
	// Timestamp    int64    `json:"ts"`
	// MarkdownIn   []string `json:"mrkdwn_in"`
	Actions    []Action `json:"actions"`
	CallbackID string   `json:"callback_id"`
	Fallback   string   `json:"fallback"`
	// ThumbnailUrl string   `json:"thumb_url"`
}

// Payload for slack
type Payload struct {
	Username     string       `json:"username,omitempty"`
	ResponseType string       `json:"response_type,omitempty"`
	IconEmoji    string       `json:"icon_emoji,omitempty"`
	Channel      string       `json:"channel,omitempty"`
	Text         string       `json:"text,omitempty"`
	LinkNames    string       `json:"link_names,omitempty"`
	Attachments  []Attachment `json:"attachments,omitempty"`
}

func question(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Slack Question endpoint")

	log.Println("Checking body")
	if r.Body == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("empty body"))
		return
	}
	defer r.Body.Close()

	log.Println("Parsing body")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("could not parse body"))
		return
	}

	log.Println("Parsing payload")
	// slack API calls the data POST a 'payload'
	log.Println(r.PostForm)

	actions := []Action{}
	actions = append(actions, Action{
		Text:  "Reveal your answer",
		Type:  "button",
		Style: "good",
		Name:  "reveal",
	})

	attachments := []Attachment{}
	attachments = append(attachments, Attachment{
		Actions:        actions,
		Fallback:       "You are unable to choose a game",
		CallbackID:     "question_of_the_day",
		Color:          "#3AA3E3",
		AttachmentType: "default",
	})
	group := Payload{
		Username:     "Quizando",
		IconEmoji:    ":bug:",
		Text:         "This is your question",
		ResponseType: "in_channel",
		Attachments:  attachments,
	}

	out, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}

	log.Println(string(out))

	res, e := http.Post(string(r.PostForm["response_url"][0]), "application/json", bytes.NewBuffer(out))

	if e != nil {
		log.Fatalln(e)
	}

	log.Println(res)
}

// Data for slack action response
type Data struct {
	TriggerID string `json:"response_url"`
}

func answer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Slack Answer endpoint")

	log.Println("Checking body")
	if r.Body == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("empty body"))
		return
	}
	defer r.Body.Close()

	log.Println("Parsing body")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("could not parse body"))
		return
	}

	log.Println("Parsing payload")
	// slack API calls the data POST a 'payload'
	reply := r.PostFormValue("payload")

	data := &Data{
		TriggerID: "",
	}
	errs := json.Unmarshal([]byte(reply), data)
	fmt.Println(errs)
	fmt.Println(data.TriggerID)

	actions := []Action{}
	actions = append(actions, Action{
		Text:  "This is my nswer",
		Type:  "mrkdwn",
		Style: "good",
		Name:  "reveal",
	})

	attachments := []Attachment{}
	attachments = append(attachments, Attachment{
		Actions:        actions,
		Fallback:       "You are unable to choose a game",
		CallbackID:     "question_of_the_day",
		Color:          "#3AA3E3",
		AttachmentType: "default",
	})
	group := Payload{
		Username:     "Quizando",
		IconEmoji:    ":bug:",
		Text:         "This is your question",
		ResponseType: "ephemeral",
		Attachments:  attachments,
	}

	out, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}

	log.Println(string(out))

	res, e := http.Post(string(data.TriggerID), "application/json", bytes.NewBuffer(out))

	if e != nil {
		log.Fatalln(e)
	}

	log.Println(res)
}
