package user

import(
	"context"
	"encoding/json"
	"net/http"

)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service: service}
}

type createUserRequest struct{
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}



func (h *Handler) Create(w http.ResponseWriter, r *http.Request){
	var request createUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != 
}