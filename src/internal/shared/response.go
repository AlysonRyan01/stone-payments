package shared

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, code int, data any, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := response{
		Data:  data,
		Error: errMsg,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}