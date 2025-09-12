package main

import (
	"fmt"
	"errors"
)

func commandExplore( c *config) error {
	if c.Args == nil || len(c.Args) == 0 {
		return errors.New("No Location name or id specified.")
	}
	data, err := c.pokeapiClient.ExploreLocation(c.Args[0])
	if err != nil {
		return err
	}
	for _, encounter := range data.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
