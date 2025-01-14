package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations

func (c *Client) ListLocations(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}
	locationResp := LocationAreas{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationAreas{}, err
	}
	return locationResp, nil
}
