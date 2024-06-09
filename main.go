package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		msg := struct {
			Message string `json:"message"`
		}{Message: "Server is alive!"}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(msg)
	}).Methods("GET")

	methods := handlers.AllowedMethods([]string{"POST"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Listening port :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(methods, ttl, origins)(api)))
}
