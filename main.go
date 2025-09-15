package main

import (
	"time"
	"github.com/jmservic/pokedexcli/internal/pokeapi"
)

func main() {
	con := &config {
		pokeapiClient: pokeapi.NewClient(5 * time.Second, 5 * time.Minute),
	}
	startRepl(con)
}
