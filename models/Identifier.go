package models

type Identifier struct {
	Use      string          `json:"use"`
	Type     *IdentifierType `json:"type"`
	System   string          `json:"system"`
	Value    string          `json:"value"`
	Period   *Period         `json:"period"`
	Assigner struct {
		Display string `json:"display"`
	} `json:"assigner"`
}

type IdentifierType struct {
	Coding []*Coding `json:"coding"`
}
