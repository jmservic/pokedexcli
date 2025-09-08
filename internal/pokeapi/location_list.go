package pokeapi

import (
	"encoding/json"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()

	var data LocationResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return LocationResponse{}, err
	}

	return data, nil
}
