package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type createUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type userResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

func newUserResponse(user *User) userResponse {
	return userResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request createUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.Create(r.Context(), CreateUserInput{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	})

	if err != nil {
    	if errors.Is(err, ErrEmailExists) {
        	http.Error(w, ErrEmailExists.Error(), http.StatusConflict)
        	return
    	}

    	http.Error(w, err.Error(), http.StatusBadRequest)
    	return
	}
	response := newUserResponse(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
