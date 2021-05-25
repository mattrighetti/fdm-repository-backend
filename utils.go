package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// sendStatus sends a status code to the client
func sendStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

// sendJSON sends a JSON message to the client
func sendJSON(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		log.Printf("error marshaling object: %v", obj)
		return
	}
	_, err = w.Write(jsonObj)
	if err != nil {
		log.Printf("error sending json response: %v", err)
		return
	}
	w.WriteHeader(code)
}
