package shared

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status        any    `json:"data,omitempty"`
	TransactionID string `json:"transactionId,omitempty"`
	Error         string `json:"error,omitempty"`
	Warning       string `json:"warning,omitempty"`
}

func WriteJSON(w http.ResponseWriter, code int, status string, data string, errMsg string, warning string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := response{
		Status: status ,
		TransactionID:  data,
		Error: errMsg,
		Warning: warning,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
