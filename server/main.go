package main

import (
	"fmt"
	"main/database"
	"main/logger"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	logger.Init()
	database.InitDB()
	defer database.DB.Close()

	r := mux.NewRouter()
	RegisterRoutes(r)
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", r)

}
