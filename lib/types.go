package lib

// OSLJSONFormat represents the JSON structure of the API for license overviews
type OSLJSONFormat struct {
	Name string `json:"name"`
	Link string `json:"link"`
	ID   string `json:"id"`
}

// Config represents the configuration for text replacements inside license templates
type Config struct {
	Name string
	Year string
}
