package transaction

import (
	"encoding/json"
	"log"
	"net/http"
	"stone/internal/shared"
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
		shared.WriteJSON(w, http.StatusBadRequest, "", "", ErrInvalidPayload.Error(), "")
		return
	}

	if err := request.Validate(); err != nil {
		shared.WriteJSON(w, http.StatusBadRequest, "", "", err.Error(), "")
		return
	}

	model, err := CreateRequestToModel(request)
	if err != nil {
		shared.WriteJSON(w, http.StatusBadRequest, "", "", err.Error(), "")
		return
	}
	
	id, err := h.service.Create(r.Context(), model)
	if err != nil {
		log.Printf("failed to create transaction: %v", err)
		shared.WriteJSON(w, http.StatusBadRequest, "", "", err.Error(), "")
		return
	}

	shared.WriteJSON(w, http.StatusCreated, map[string]string{"id": id}, "")
}