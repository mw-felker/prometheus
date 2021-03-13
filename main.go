package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/mw-felker/prometheus/models"
)

const resourceBaseURL = "https://www.hl7.org/fhir"

// https://github.com/samply/golang-fhir-models/blob/master/fhir-models/gen-resources.sh

func main() {
	newPatient := mapPatientJSONtoStruct()
	fmt.Println("Mapped " + newPatient.ResourceType + " JSON to struct: " + newPatient.Gender + " " + newPatient.BirthDate + " from " + newPatient.Address[0].City)
	//	encounterExample := string(getEncounterJSON())
	//fmt.Println("Retrieved Encounter shape from HL7 site")
	//fmt.Println(encounterExample)
}

func mapPatientJSONtoStruct() models.Patient {
	var newPatient models.Patient
	patientJSON := getPatientJSON()
	error := json.Unmarshal(patientJSON, &newPatient)
	handleError(error)
	return newPatient
}

func getEncounterJSON() []byte {
	const url string = resourceBaseURL + "/encounter-example.json"
	return retrieveJSONData(url)
}

func getPatientJSON() []byte {
	const url string = resourceBaseURL + "/patient-example.json"
	return retrieveJSONData(url)
}

func retrieveJSONData(url string) []byte {
	response, error := http.Get(url)
	handleError(error)
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	handleError(error)
	return body
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
