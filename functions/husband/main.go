package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	apex "github.com/apex/go-apex"
)

type AlexaRequest struct {
	Session struct {
		SessionID   string `json:"sessionId"`
		Application struct {
			ApplicationID string `json:"applicationId"`
		} `json:"application"`
		Attributes struct {
		} `json:"attributes"`
		User struct {
			UserID string `json:"userId"`
		} `json:"user"`
		New bool `json:"new"`
	} `json:"session"`
	Request struct {
		Type      string    `json:"type"`
		RequestID string    `json:"requestId"`
		Locale    string    `json:"locale"`
		Timestamp time.Time `json:"timestamp"`
		Intent    struct {
			Name  string `json:"name"`
			Slots struct {
			} `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
	Version string `json:"version"`
}

type AlexaReply struct {
	Version           string       `json:"version"`
	Response          ResponseType `json:"response"`
	SessionAttributes struct {
	} `json:"sessionAttributes"`
}

type ResponseType struct {
	OutputSpeech     OutputSpeechType `json:"outputSpeech"`
	Card             CardType         `json:"card"`
	ShouldEndSession bool             `json:"shouldEndSession"`
}

type OutputSpeechType struct {
	Type string `json:"type"`
	Ssml string `json:"ssml"`
}

type CardType struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

type SessionAttributesType struct{}

func reply(text string) AlexaReply {
	return AlexaReply{
		Version: "1.0",
		Response: ResponseType{
			OutputSpeech: OutputSpeechType{
				Type: "SSML",
				Ssml: "<speak> Info: The tapir’s nose is prehensile and is used to grab leaves – and also as a snorkel while swimming. </speak>",
			},
			Card: CardType{
				Content: "The tapir’s nose is prehensile and is used to grab leaves – and also as a snorkel while swimming.",
				Title:   "Tapir Info",
				Type:    "Simple",
			},
			ShouldEndSession: true,
		},
		SessionAttributes: SessionAttributesType{},
	}
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var request AlexaRequest

		if err := json.Unmarshal(event, &request); err != nil {
			return nil, err
		}
		fmt.Fprintf(os.Stderr, "request: %v\n", request)
		fmt.Fprintf(os.Stderr, "request: %#v\n", request)

		return reply("dingo"), nil
	})
}
