package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type awsJson struct {
	Message string `json:"Message"`
}

type message struct {
	EndpointArn      string `json:"EndpointArn"`
	DeliveryAttempts int    `json:"DeliveryAttempts"`
}

func fail_send(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	json_message := MessageFromAWSJson(body)

	if json_message.DeliveryAttempts > 2 {
		pusher.DelDeviceWithArn(json_message.EndpointArn)
		log.Println("Removed a device. ARN: ", json_message.EndpointArn)
	}

	success(&w, Response{})
}

func MessageFromAWSJson(j []byte) (m message) {
	var json_struct awsJson

	json.Unmarshal(j, &json_struct)
	json.Unmarshal([]byte(json_struct.Message), &m)

	return
}
