package organization

import (
	"encoding/json"
	"net/http"
)

type Handler struct{
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service: service}
}
type createOrganizationRequest struct{
	Name string `json:"name"`
}


func (h *Handler) Create(w http.ResponseWriter, r *http.Request){
	var request createOrganizationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	organization, err := h.service.Create(r.Context(), request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(organization)
}