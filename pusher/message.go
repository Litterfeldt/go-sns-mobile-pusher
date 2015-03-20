package pusher

import (
	"encoding/json"
	"strconv"
)

type Message struct {
	PushToken   string
	Text        string
	Url         string
	UnreadCount string
}

type JsonMessage struct {
	Default string `json:"default"`
	GCM     string `json:"GCM"`
	APNS    string `json:"APNS"`
}

type AndroidMessage struct {
	Data AndroidData `json:"data"`
}

type AndroidData struct {
	Message     string `json:"message"`
	UnreadCount string `json:"unread_notification_count"`
	Url         string `json:"url"`
}

type AppleMessage struct {
	Aps AppleData `json:"aps"`
}

type AppleData struct {
	Alert string `json:"alert"`
	Badge int    `json:"badge"`
	Url   string `json:"url"`
}

func (m *Message) ToJson() string {
	json_h := JsonMessage{
		//Default: m.Text,
		GCM:  androidJSON(m),
		APNS: appleJSON(m),
	}

	json, _ := json.Marshal(json_h)
	return string(json)
}

func appleJSON(m *Message) string {
	int_count, _ := strconv.Atoi(m.UnreadCount)
	apple_h := AppleMessage{
		Aps: AppleData{
			Alert: m.Text,
			Badge: int_count,
			Url:   m.Url,
		},
	}
	apple_json, err := json.Marshal(apple_h)
	if err != nil {
		return ""
	}
	return string(apple_json)
}
func androidJSON(m *Message) string {
	android_h := AndroidMessage{
		Data: AndroidData{
			Message:     m.Text,
			UnreadCount: m.UnreadCount,
			Url:         m.Url,
		},
	}
	android_json, err := json.Marshal(android_h)
	if err != nil {
		return ""
	}
	return string(android_json)
}
