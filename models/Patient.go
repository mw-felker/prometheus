package models

type Patient struct {
	ResourceType         string                `json:"resourceType"`
	ID                   string                `json:"id"`
	Text                 *Text                 `json:"text"`
	Identifier           []*Identifier         `json:"identifier"`
	Active               bool                  `json:"active"`
	Name                 []*Name               `json:"name"`
	Telecom              []*Telecom            `json:"telecom"`
	Gender               string                `json:"gender"`
	BirthDate            string                `json:"birthDate"`
	UnderscoreBirthDate  *BirthDateExtenstion  `json:"_birthDate"`
	DeceasedBoolean      bool                  `json:"deceasedBoolean"`
	Address              []*Address            `json:"address"`
	Contact              []*Contact            `json:"contact"`
	ManagingOrganization *ManagingOrganization `json:"managingOrganization"`
}
