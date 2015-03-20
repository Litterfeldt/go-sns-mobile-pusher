package api

import (
	p "github.com/litterfeldt/go-sns-mobile-pusher/pusher"
	"log"
	"net/http"
	"os"
)

type Response map[string]interface{}

var pusher *p.Pusher

func Start() {
	pusher = p.New()
	http.HandleFunc("/status", auth(status_handler))
	http.HandleFunc("/send", auth(send_handler))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Auth-Token") == os.Getenv("AUTH_TOKEN") {
			h.ServeHTTP(w, r)
			return
		}
		unauthorized(&w)
	}
}

func status_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		success(&w, Response{
			"running":       true,
			"workers":       len(p.WorkerQueue),
			"jobs_in_queue": len(p.WorkQueue),
		})
	} else {
		not_found(&w, nil)
	}
}

func send_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		send(w, r)
	} else {
		not_found(&w, nil)
	}
}
