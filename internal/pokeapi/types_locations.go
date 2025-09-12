package pokeapi

// locationResponse -
type LocationResponse struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

type Location struct {
	Id int `json:"id"`
	Name string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
