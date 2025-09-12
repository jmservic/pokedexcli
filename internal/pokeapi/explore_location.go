package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

//ExploreLocations -
func (c *Client) ExploreLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location
	
	rawData, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Location{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		}

		rawData, err = io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}

		c.cache.Add(url, rawData)
	}
	data := Location{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return Location{}, err
	}
	return data, nil

}
