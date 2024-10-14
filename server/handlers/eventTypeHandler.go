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

func CreateEventType(w http.ResponseWriter, r *http.Request) {
	var eventType models.EventType
	if err := json.NewDecoder(r.Body).Decode(&eventType); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := eventType.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventType.SetDefaultValues()

	if err := database.CreateEventType(eventType); err != nil {
		log.Errorf("Failed to create event type: %v", err)
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1062:
				http.Error(w, "Event type with this name already exists", http.StatusConflict)
				return
			}

		}
		http.Error(w, "Failed to create event type", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(eventType)
}

func GetEventTypes(w http.ResponseWriter, r *http.Request) {
	eventTypes, err := database.GetAllEventTypes()
	if err != nil {
		log.Errorf("Failed to get event types: %v", err)
		http.Error(w, "Failed to get event types", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(eventTypes)
}

func GetEventTypeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	eventType, err := database.GetEventTypeByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Event type not found", http.StatusNotFound)
			return
		}
		log.Errorf("Failed to get event type: %v", err)
		http.Error(w, "Failed to get event type", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(eventType)
}

func DeleteEventType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the event type exists
	if _, err := database.GetEventTypeByID(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Event type not found", http.StatusNotFound)
			return
		}
		log.Errorf("Failed to get event type: %v", err)
		http.Error(w, "Failed to get event type", http.StatusInternalServerError)
		return
	}

	// Delete the event type
	if err := database.DeleteEventType(id); err != nil {
		log.Errorf("Failed to delete event type: %v", err)
		http.Error(w, "Failed to delete event type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PutEventTypeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid event type id", http.StatusBadRequest)
		return
	}

	var eventType models.EventType
	if err := json.NewDecoder(r.Body).Decode(&eventType); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := eventType.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.PutEventTypeByID(id, eventType); err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1062:
				http.Error(w, "Event type with this name already exists", http.StatusConflict)
				return
			}
		}
		log.Errorf("Failed to update event type: %v", err)
		http.Error(w, "Failed to update event type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
