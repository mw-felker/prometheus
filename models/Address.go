package models

type Address struct {
	Use        string   `json:"use"`
	Type       string   `json:"type"`
	Text       string   `json:"text"`
	Line       []string `json:"line"`
	City       string   `json:"city"`
	District   string   `json:"district"`
	State      string   `json:"state"`
	PostalCode string   `json:"postalCode"`
	Period     *Period  `json:"period"`
}
