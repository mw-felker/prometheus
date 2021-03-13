package models

type Contact struct {
	Relationship []*Relationship `json:"relationship"`
	Name         *Name           `json:"name"`
	Telecom      []*Telecom      `json:"telecom"`
	Address      *Address        `json:"address"`
	Gender       string          `json:"gender"`
	Period       *Period         `json:"period"`
}
