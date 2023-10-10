package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Start the RestAPI
func Run() {
	r := mux.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                                       // Set the allowed origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Set the allowed methods
		handlers.AllowedHeaders([]string{"Content-Type"}),                            // Set the allowed headers
	)

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/api/applications", getApplications).Methods("GET")
	r.HandleFunc("/api/application/{name}", getApplicaton).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
