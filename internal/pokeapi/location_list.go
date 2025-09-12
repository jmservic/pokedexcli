package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	//"fmt"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	rawData, ok := c.cache.Get(url)

	if !ok {
	//	fmt.Println("Did not use the cache\n")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationResponse{}, err
		}
		defer res.Body.Close()
		
		rawData, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationResponse{}, err
		}
		c.cache.Add(url, rawData)
	}/* else {
		fmt.Println("Hmm... the cache came in handy.\n")
	}*/


	data := LocationResponse{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return data, err
	} 
	return data, nil
}
