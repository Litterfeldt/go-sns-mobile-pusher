package api

import (
	"encoding/json"
	p "github.com/litterfeldt/go-sns-mobile-pusher/pusher"
	"log"
	"net/http"
)

type Message struct {
	PushToken   string `json:"push_token"`
	Message     string `json:"message"`
	Url         string `json:"url"`
	UnreadCount string `json:"unread_count"`
}

func send(w http.ResponseWriter, r *http.Request) {
	m, err := decodeMessage(r)
	if err != nil {
		log.Println(err)
		client_error(&w)
		return
	}

	p.WorkQueue <- p.Message{
		PushToken:   m.PushToken,
		Text:        m.Message,
		Url:         m.Url,
		UnreadCount: m.UnreadCount,
	}

	success(&w, Response{})
}

func decodeMessage(r *http.Request) (m Message, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&m)
	return
}
