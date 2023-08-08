package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Something went wrong while marshalling the payload to JSON\n payload : %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, errStr string) {
	respondWithJson(w, code, errorResponse{Error: errStr})
}

type errorResponse struct {
	Error string `json:"error"`
}
