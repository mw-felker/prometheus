package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

	getPatientShape()

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

func getPatientShape() {
	const url string = "https://www.hl7.org/fhir/patient-example.json"
	const filepath string = "Patient.json"

	error := DownloadFile(url, filepath)
	handleError(error)

	patientShape, error := ioutil.ReadFile(filepath)
	handleError(error)

	fmt.Println("Reading local file " + filepath)
	fmt.Print(string(patientShape))
}

func DownloadFile(url string, filepath string) error {

	// Get the data
	response, error := http.Get(url)
	handleError(error)
	defer response.Body.Close()

	fmt.Println("Successfully downloaded: " + url)

	// Create the file
	newFile, error := os.Create(filepath)
	handleError(error)
	defer newFile.Close()

	fmt.Println("Successfully created file " + filepath)

	// Write the body to file
	_, error = io.Copy(newFile, response.Body)
	handleError(error)

	return nil
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
