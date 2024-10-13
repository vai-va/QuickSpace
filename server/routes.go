package main

import (
	"main/handlers"
	"main/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/user", middleware.Logging(handlers.CreateUser)).Methods("POST")
	router.HandleFunc("/api/users", middleware.Logging(handlers.GetUsers)).Methods("GET")
	router.HandleFunc("/api/user/{id}", middleware.Logging(handlers.GetUserByID)).Methods("GET")
	router.HandleFunc("/api/user/{id}", middleware.Logging(handlers.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/api/user/{user_id}/space_reservations", middleware.Logging(handlers.GetSpaceReservationsByRentedById)).Methods("GET")

	router.HandleFunc("/api/space", middleware.Logging(handlers.CreateSpace)).Methods("POST")
	router.HandleFunc("/api/spaces", middleware.Logging(handlers.GetSpaces)).Methods("GET")
	router.HandleFunc("/api/space/{id}", middleware.Logging(handlers.GetSpaceByID)).Methods("GET")
	router.HandleFunc("/api/space/{id}", middleware.Logging(handlers.DeleteSpace)).Methods("DELETE")
	router.HandleFunc("/api/space/{space_id}/ratings", middleware.Logging(handlers.GetSpaceRatingsBySpaceID)).Methods("GET")

	router.HandleFunc("/api/event_type", middleware.Logging(handlers.CreateEventType)).Methods("POST")
	router.HandleFunc("/api/event_types", middleware.Logging(handlers.GetEventTypes)).Methods("GET")
	router.HandleFunc("/api/event_type/{id}", middleware.Logging(handlers.GetEventTypeByID)).Methods("GET")
	router.HandleFunc("/api/event_type/{id}", middleware.Logging(handlers.DeleteEventType)).Methods("DELETE")

	router.HandleFunc("/api/space_reservation", middleware.Logging(handlers.CreateSpaceReservation)).Methods("POST")
	router.HandleFunc("/api/space_reservations", middleware.Logging(handlers.GetSpaceReservations)).Methods("GET")
	router.HandleFunc("/api/space_reservation/{id}", middleware.Logging(handlers.GetSpaceReservationByID)).Methods("GET")
	router.HandleFunc("/api/space_reservation/{id}/status", middleware.Logging(handlers.UpdateSpaceReservationStatus)).Methods("PATCH")

	router.HandleFunc("/api/space_rating", middleware.Logging(handlers.CreateSpaceRating)).Methods("POST")
	router.HandleFunc("/api/space_ratings", middleware.Logging(handlers.GetSpaceRatings)).Methods("GET")
	router.HandleFunc("/api/space_rating/{rating_id}", middleware.Logging(handlers.GetSpaceRatingByID)).Methods("GET")
	router.HandleFunc("/api/space_rating/{rating_id}", middleware.Logging(handlers.DeleteSpaceRating)).Methods("DELETE")

	// Get all of specific user's rented spaces based on event type they rented for
	router.HandleFunc("/api/user/{user_id}/reservation_event_type/{event_type_id}/spaces", middleware.Logging(handlers.GetSpacesByUserIDByEventType)).Methods("GET")
}
