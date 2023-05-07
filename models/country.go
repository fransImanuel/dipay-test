package models

type Country struct {
	Name     string   `json:"name"`
	Region   string   `json:"region"`
	Timezone []string `json:"timezones"`
}

type CountryCollection struct {
	Result []Country
}
