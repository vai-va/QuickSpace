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

func CreateSpaceRating(w http.ResponseWriter, r *http.Request) {
	var spaceRating models.SpaceRating
	if err := json.NewDecoder(r.Body).Decode(&spaceRating); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := spaceRating.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	spaceRating.SetDefaultValues()

	if err := database.CreateSpaceRating(spaceRating); err != nil {
		log.Errorf("Failed to create space rating: %v", err)
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			log.Errorf("Failed to create space rating: %v", sqlErr)
			http.Error(w, "Failed to create space rating", http.StatusInternalServerError)

		}
		http.Error(w, "Failed to create space rating", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(spaceRating)
}

func GetSpaceRatingsBySpaceID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["space_id"]

	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid space ID", http.StatusBadRequest)
		return
	}

	spaceRatings, err := database.GetSpaceRatingsBySpaceID(id)
	if err != nil {
		log.Errorf("Failed to get space ratings: %v", err)
		http.Error(w, "Failed to get space ratings", http.StatusInternalServerError)
		return
	}

	if len(spaceRatings) == 0 {
		http.Error(w, "No ratings found for the space", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(spaceRatings)
}

func GetSpaceRatingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["rating_id"]
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid rating ID", http.StatusBadRequest)
		return
	}

	spaceRating, err := database.GetSpaceRatingByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Rating not found", http.StatusNotFound)
			return
		}
		log.Errorf("Failed to get space rating: %v", err)
		http.Error(w, "Failed to get space rating", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(spaceRating)
}

func DeleteSpaceRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["rating_id"]
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid rating ID", http.StatusBadRequest)
		return
	}

	// Check if the rating exists
	if _, err := database.GetSpaceRatingByID(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Rating not found", http.StatusNotFound)
			return
		}
		log.Errorf("Failed to get rating: %v", err)
		http.Error(w, "Failed to get rating", http.StatusInternalServerError)
		return
	}

	// Delete the rating
	if err := database.DeleteSpaceRating(id); err != nil {
		log.Errorf("Failed to delete space rating: %v", err)
		http.Error(w, "Failed to delete space rating", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetSpaceRatings(w http.ResponseWriter, r *http.Request) {
	spaceRatings, err := database.GetAllSpaceRatings()
	if err != nil {
		log.Errorf("Failed to get space ratings: %v", err)
		http.Error(w, "Failed to get space ratings", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(spaceRatings)
}
