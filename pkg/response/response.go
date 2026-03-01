package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("json encode error: %v\n", err)
	}
}

func Error(w http.ResponseWriter, status int, msg string) {
	type errResponse struct {
		Error string `json:"error"`
	}
	JSON(w, status, errResponse{Error: msg})
}
