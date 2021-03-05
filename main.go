package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	mapPatientJSONtoStruct()
	fmt.Println(string(getEncounterJSON()))
}

/*
func serveAPI() {
	const port = ":4200"
	router := mux.NewRouter().StrictSlash(true)

	//patients = append(patients, Patient{UUID: uuid.NewString(), FirstName: "Max", LastName: "Felker"})
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
}*/

func mapPatientJSONtoStruct() {
	var newPatient Patient
	patientJSON := getPatientJSON()
	error := json.Unmarshal(patientJSON, &newPatient)
	handleError(error)
	fmt.Println(newPatient.ResourceType + " created: " + newPatient.Gender + " " + newPatient.BirthDate)
}

func getEncounterJSON() []byte {
	const url string = "https://www.hl7.org/fhir/encounter-example.json"
	return retrieveJSONData(url)
}

func getPatientJSON() []byte {
	const url string = "https://www.hl7.org/fhir/patient-example.json"
	return retrieveJSONData(url)
}

func retrieveJSONData(url string) []byte {
	response, error := http.Get(url)
	handleError(error)
	defer response.Body.Close()

	fmt.Println("Successfully retrieved " + url)

	body, error := ioutil.ReadAll(response.Body)
	handleError(error)
	return body
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

// Patient FHIR struct
type Patient struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Text         struct {
		Status string `json:"status"`
		Div    string `json:"div"`
	} `json:"text"`
	Identifier []struct {
		Use  string `json:"use"`
		Type struct {
			Coding []struct {
				System string `json:"system"`
				Code   string `json:"code"`
			} `json:"coding"`
		} `json:"type"`
		System string `json:"system"`
		Value  string `json:"value"`
		Period struct {
			Start string `json:"start"`
		} `json:"period"`
		Assigner struct {
			Display string `json:"display"`
		} `json:"assigner"`
	} `json:"identifier"`
	Active bool `json:"active"`
	Name   []struct {
		Use    string   `json:"use"`
		Family string   `json:"family,omitempty"`
		Given  []string `json:"given"`
		Period struct {
			End string `json:"end"`
		} `json:"period,omitempty"`
	} `json:"name"`
	Telecom []struct {
		Use    string `json:"use"`
		System string `json:"system,omitempty"`
		Value  string `json:"value,omitempty"`
		Rank   int    `json:"rank,omitempty"`
		Period struct {
			End string `json:"end"`
		} `json:"period,omitempty"`
	} `json:"telecom"`
	Gender              string `json:"gender"`
	BirthDate           string `json:"birthDate"`
	UnderscoreBirthDate struct {
		Extension []struct {
			URL           string `json:"url"`
			ValueDateTime string `json:"valueDateTime"`
		} `json:"extension"`
	} `json:"_birthDate"`
	DeceasedBoolean bool `json:"deceasedBoolean"`
	Address         []struct {
		Use        string   `json:"use"`
		Type       string   `json:"type"`
		Text       string   `json:"text"`
		Line       []string `json:"line"`
		City       string   `json:"city"`
		District   string   `json:"district"`
		State      string   `json:"state"`
		PostalCode string   `json:"postalCode"`
		Period     struct {
			Start string `json:"start"`
		} `json:"period"`
	} `json:"address"`
	Contact []struct {
		Relationship []struct {
			Coding []struct {
				System string `json:"system"`
				Code   string `json:"code"`
			} `json:"coding"`
		} `json:"relationship"`
		Name struct {
			Family           string `json:"family"`
			UnderscoreFamily struct {
				Extension []struct {
					URL         string `json:"url"`
					ValueString string `json:"valueString"`
				} `json:"extension"`
			} `json:"_family"`
			Given []string `json:"given"`
		} `json:"name"`
		Telecom []struct {
			System string `json:"system"`
			Value  string `json:"value"`
		} `json:"telecom"`
		Address struct {
			Use        string   `json:"use"`
			Type       string   `json:"type"`
			Line       []string `json:"line"`
			City       string   `json:"city"`
			District   string   `json:"district"`
			State      string   `json:"state"`
			PostalCode string   `json:"postalCode"`
			Period     struct {
				Start string `json:"start"`
			} `json:"period"`
		} `json:"address"`
		Gender string `json:"gender"`
		Period struct {
			Start string `json:"start"`
		} `json:"period"`
	} `json:"contact"`
	ManagingOrganization struct {
		Reference string `json:"reference"`
	} `json:"managingOrganization"`
}
