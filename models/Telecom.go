package models

type Telecom struct {
	Use    string  `json:"use"`
	System string  `json:"system,omitempty"`
	Value  string  `json:"value,omitempty"`
	Rank   int     `json:"rank,omitempty"`
	Period *Period `json:"period,omitempty"`
}
