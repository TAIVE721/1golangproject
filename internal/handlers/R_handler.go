package handlers

import (
	"RiderApi/internal/domain"
	"RiderApi/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RiderHandler struct {
	service services.RiderService
}

func NewRiderHandler(s services.RiderService) *RiderHandler {
	return &RiderHandler{
		service: s,
	}
}

func (r *RiderHandler) GetAll(w http.ResponseWriter, h *http.Request) {

	riders, err := r.service.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(riders)

}

func (r *RiderHandler) GetById(w http.ResponseWriter, h *http.Request) {
	id := chi.URLParam(h, "Rider_id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, ("Formato de Id invalido "), http.StatusBadRequest)
		return
	}

	rider, err := r.service.GetById(ID)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rider)

}

func (r *RiderHandler) Post(w http.ResponseWriter, h *http.Request) {

	var rider domain.KamenRider

	if err := json.NewDecoder(h.Body).Decode(&rider); err != nil {

		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	Rider, err := r.service.Post(rider)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Rider)

}

func (r *RiderHandler) Patch(w http.ResponseWriter, h *http.Request) {

	id := chi.URLParam(h, "Rider_id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Formato de id incompatible", http.StatusBadRequest)
		return
	}

	var rider domain.KamenRider

	if err := json.NewDecoder(h.Body).Decode(&rider); err != nil {
		http.Error(w, "Formato de id incompatible", http.StatusBadRequest)
		return
	}

	PatchedRider, err := r.service.Patch(rider, ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PatchedRider)

}

func (r *RiderHandler) Delete(w http.ResponseWriter, h *http.Request) {

	id := chi.URLParam(h, "Rider_id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, ("Formato de id incompatible"), http.StatusBadRequest)
		return
	}

	LastId, err := r.service.Delete(ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(LastId)

}
