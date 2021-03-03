package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Patient struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

const port = ":4200"

var patients []Patient

func main() {
	router := mux.NewRouter().StrictSlash(true)
	patients = append(patients, Patient{UUID: uuid.NewString(), FirstName: "Max", LastName: "Felker"})
	registerRoutes(router)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(port, router))
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/Patients", getPatientHandler).Methods("GET")
}

func getPatientHandler(writer http.ResponseWriter, response *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(patients)
}
