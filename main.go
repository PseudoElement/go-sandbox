package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pseudoelement/go-sandbox/common/constants"
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

	api.HandleFunc("/tonconnect/manifest/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		path, ok := constants.URL_TO_MANIFEST_PATH_MAP[id]
		if !ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Invalid environvent id!"))
			return
		}
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, path)
	})

	api.HandleFunc("/tonconnect/logo.png", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, "./static/rubic-logo.png")
	})

	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Listening port :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(methods, ttl, origins)(api)))
}
