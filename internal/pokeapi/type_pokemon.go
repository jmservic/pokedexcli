package pokeapi


type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BaseExp int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort int `json:"effort"`
		Stat struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"stat"`
	} `json:"stats"` 
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
