package transaction

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, ErrInvalidPayload.Error())
		return
	}

	if err := request.Validate(); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	model, err := CreateRequestToModel(request)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	
	id, err := h.service.Create(r.Context(), model)
	if err != nil {
		log.Printf("failed to create transaction: %v", err)
		writeJSON(w, http.StatusInternalServerError, nil, "internal server error")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"id": id}, "")
}

type response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, code int, data any, errMsg string) {
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