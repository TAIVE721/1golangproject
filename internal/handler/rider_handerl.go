package handler

import (
	"RiderApi/internal/service"
	"database/sql"
	"encoding/json"
	"net/http"
)

type RiderHandler struct {
	service service.RiderService
}

func NewRiderHandler(db *sql.DB) *RiderHandler {
	return &RiderHandler{
		service: service.NewRiderService(db),
	}
}

func (h *RiderHandler) GetAllRiders(w http.ResponseWriter, r *http.Request) {

	riders, err := h.service.GetAllRiders()

	if err != nil {
		http.Error(w, "Error al obtener los riders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/json")

	json.NewEncoder(w).Encode(riders)

}
