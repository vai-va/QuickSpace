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

func CreateSpace(w http.ResponseWriter, r *http.Request) {
	var space models.Space
	if err := json.NewDecoder(r.Body).Decode(&space); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := space.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	space.SetDefaultValues()

	if err := database.CreateSpace(space); err != nil {
		log.Errorf("Failed to create space: %v", err)
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			log.Errorf("Failed to create space: %v", sqlErr)
			http.Error(w, "Failed to create space", http.StatusInternalServerError)

		}
		http.Error(w, "Failed to create space", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(space)
}

func GetSpaces(w http.ResponseWriter, r *http.Request) {
	spaces, err := database.GetAllSpaces()
	if err != nil {
		log.Errorf("Failed to get spaces: %v", err)
		http.Error(w, "Failed to get spaces", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(spaces)
}

func GetSpaceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	space, err := database.GetSpaceByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Space not found", http.StatusNotFound)
			return
		}
		log.Errorf("Failed to get space: %v", err)
		http.Error(w, "Failed to get space", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(space)
}

func DeleteSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the space exists
	_, err := database.GetSpaceByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Space not found", http.StatusNotFound)
			return
		}
		log.Errorf("Failed to get user: %v", err)
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// Delete the space
	err = database.DeleteSpace(id)
	if err != nil {
		log.Errorf("Failed to delete space: %v", err)
		http.Error(w, "Failed to delete space", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetSpacesByUserIDByEventType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	eventType_id := vars["event_type_id"]
	if _, err := uuid.Parse(userID); err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if _, err := uuid.Parse(eventType_id); err != nil {
		http.Error(w, "Invalid event type id", http.StatusBadRequest)
		return
	}
	spaces, err := database.GetSpacesByUserIDByEventType(userID, eventType_id)
	if err != nil {
		log.Errorf("Failed to get spaces: %v", err)
		http.Error(w, "Failed to get spaces", http.StatusInternalServerError)
		return
	}

	if len(spaces) == 0 {
		http.Error(w, "No spaces found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(spaces)
}
