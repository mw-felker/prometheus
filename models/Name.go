package models

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
