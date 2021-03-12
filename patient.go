package main

type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Text struct {
	Status string `json:"status"`
	Div    string `json:"div"`
}

type Coding struct {
	System string `json:"system"`
	Code   string `json:"code"`
}

type IdentifierType struct {
	Coding []*Coding `json:"coding"`
}

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

type Name struct {
	Use              string   `json:"use"`
	Family           string   `json:"family,omitempty"`
	Given            []string `json:"given"`
	Period           *Period  `json:"period,omitempty"`
	UnderscoreFamily struct {
		Extension []struct {
			URL         string `json:"url"`
			ValueString string `json:"valueString"`
		} `json:"extension"`
	} `json:"_family"`
}

type Telecom struct {
	Use    string  `json:"use"`
	System string  `json:"system,omitempty"`
	Value  string  `json:"value,omitempty"`
	Rank   int     `json:"rank,omitempty"`
	Period *Period `json:"period,omitempty"`
}

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

type Relationship struct {
	Coding []*Coding `json:"coding"`
}

type Contact struct {
	Relationship []*Relationship `json:"relationship"`
	Name         *Name           `json:"name"`
	Telecom      []*Telecom      `json:"telecom"`
	Address      *Address        `json:"address"`
	Gender       string          `json:"gender"`
	Period       *Period         `json:"period"`
}

type BirthDateExtenstion struct {
	Extension []struct {
		URL           string `json:"url"`
		ValueDateTime string `json:"valueDateTime"`
	} `json:"extension"`
}

type ManagingOrganization struct {
	Reference string `json:"reference"`
}

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
