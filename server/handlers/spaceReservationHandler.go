package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"main/database"
	"main/models"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func CreateSpaceReservation(w http.ResponseWriter, r *http.Request) {
	var spaceReservation models.SpaceReservation
	if err := json.NewDecoder(r.Body).Decode(&spaceReservation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := spaceReservation.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	spaceReservation.SetDefaultValues()

	if err := database.CreateSpaceReservation(spaceReservation); err != nil {
		log.Errorf("Failed to create space reservation: %v", err)
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			log.Errorf("Failed to create space reservation: %v", sqlErr)
			http.Error(w, "Failed to create space reservation", http.StatusInternalServerError)
		}
		http.Error(w, "Failed to create space reservation", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(spaceReservation)
}

func GetSpaceReservations(w http.ResponseWriter, r *http.Request) {
	spaceReservations, err := database.GetAllSpaceReservations()
	if err != nil {
		log.Errorf("Failed to get space reservations: %v", err)
		http.Error(w, "Failed to get space reservations", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(spaceReservations)
}

func GetSpaceReservationsByRentedById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["user_id"]
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	spaceReservations, err := database.GetSpaceReservationsByRentedById(id)
	if err != nil {
		log.Errorf("Failed to get space reservations: %v", err)
		http.Error(w, "Failed to get space reservations", http.StatusInternalServerError)
		return
	}

	if len(spaceReservations) == 0 {
		http.Error(w, "No space reservations found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(spaceReservations)
}

func GetSpaceReservationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	spaceReservation, err := database.GetSpaceReservationByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Reservation not found", http.StatusNotFound)
			return
		}
		log.Errorf("Failed to get reservation: %v", err)
		http.Error(w, "Failed to get reservation", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(spaceReservation)
}

func UpdateSpaceReservationStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updateStatusRequest models.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&updateStatusRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := updateStatusRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.UpdateSpaceReservationStatus(id, updateStatusRequest.Status); err != nil {
		log.Errorf("Failed to update space reservation status: %v", err)
		http.Error(w, "Failed to update space reservation status", http.StatusInternalServerError)
		return
	}

	spaceReservation, err := database.GetSpaceReservationByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Reservation not found", http.StatusNotFound)
			return
		}
		log.Errorf("Failed to get reservation: %v", err)
		http.Error(w, "Failed to get reservation", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(spaceReservation)
}
