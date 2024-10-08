package main

import (
	"github.com/alex-bogatiuk/integris/internal/config"
	"github.com/alex-bogatiuk/integris/internal/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config.InitConfig()
	database.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
