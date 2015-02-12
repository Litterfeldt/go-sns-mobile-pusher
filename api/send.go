package api

import (
	"encoding/json"
	p "github.com/VideofyMe/go-push-handler/pusher"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Message     string `json:"message"`
	Url         string `json:"url"`
	UserId      string `json:"user_id"`
	UnreadCount string `json:"unread_count"`
}

func send(w http.ResponseWriter, r *http.Request) {
	m, err := decodeMessage(r)
	if err != nil {
		log.Println(err)
		client_error(&w)
		return
	}

	time.Sleep(10 * time.Millisecond)

	p.WorkQueue <- p.Message{
		UserId:      m.UserId,
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
