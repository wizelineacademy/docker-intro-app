package utils

import (
	"encoding/json"
	"net/http"
)

type MainResponse struct {
	Status   int
	Message  interface{}
	Response interface{}
}

func WriteJson(w http.ResponseWriter, response interface{}, status int) {
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
