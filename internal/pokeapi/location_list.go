package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ListLocations

func (c *Client) ListLocations(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check the cache first
	if val, ok := c.cache.Get(url); ok {
		log.Printf("[CACHE] Cache hit for URL: %s\n", url)
		locationResp := LocationAreas{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationAreas{}, err
		}
		return locationResp, nil
	}

	// Log a cache miss
	log.Printf("[CACHE] Cache miss for URL: %s, fetching from API...\n", url)

	// Make an HTTP request if not found in cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	// Unmarshal the data into the response object
	locationResp := LocationAreas{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationAreas{}, err
	}
	// Log the new addition to the cache
	log.Printf("[CACHE] Adding data to cache for URL: %s\n", url)
	c.cache.Add(url, data)
	return locationResp, nil
}
