package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

//PokemonInfo -
func (c *Client) PokemonInfo(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	
	rawData, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err 
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		rawData, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, rawData)
	}

	data := Pokemon{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return Pokemon{}, err 
	}

	return data, nil
}
