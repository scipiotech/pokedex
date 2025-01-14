package pokeapi

// locations struct

type LocationAreas struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []areaData `json:"results"`
}

type areaData struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
