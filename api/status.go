package api

import (
	"encoding/json"
	"net/http"
)

func client_error(res *http.ResponseWriter) {
	http.Error(*res, "Client Error", 400)
}

func not_found(res *http.ResponseWriter, obj interface{}) {
	json_str, err := json.Marshal(obj)
	if err != nil {
		json_str = []byte("Not Found")
	}
	http.Error(*res, string(json_str), 404)
}

func server_error(res *http.ResponseWriter) {
	http.Error(*res, "Server Error", 500)
}

func success(res *http.ResponseWriter, obj interface{}) {
	json_str, _ := json.Marshal(obj)
	http.Error(*res, string(json_str), 200)
}

func unauthorized(res *http.ResponseWriter) {
	http.Error(*res, "Unauthorized", 401)
}
