package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Device struct {
	UserId      string `json:"user_id"`
	PushToken   string `json:"push_token"`
	DeviceBrand string `json:"device_brand"`
}

func create_device(w http.ResponseWriter, r *http.Request) {
	d, err := decodeDevice(r)
	if err != nil {
		log.Println(err)
		client_error(&w)
		return
	}
	device, err := pusher.AddDevice(d.UserId, d.PushToken, d.DeviceBrand)
	if err != nil {
		log.Println(err)
		server_error(&w)
		return
	}
	success(&w, Response{"success": true, "device": device})
}

func delete_device(w http.ResponseWriter, r *http.Request) {
	d, err := decodeDevice(r)
	if err != nil {
		client_error(&w)
		return
	}
	ok, _ := pusher.DelDevice(d.PushToken)
	if ok {
		success(&w, Response{"success": ok})
	} else {
		not_found(&w, Response{"success": ok})
	}
}

func decodeDevice(r *http.Request) (d Device, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&d)
	return
}
