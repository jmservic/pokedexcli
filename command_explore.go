package main

import (
	"fmt"
	"errors"
)

func commandExplore( c *config) error {
	if c.Args == nil || len(c.Args) == 0 {
		return errors.New("No Location name or id specified.")
	}
	fmt.Printf("Exploring %v...\n", c.Args[0])
	data, err := c.pokeapiClient.ExploreLocation(c.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon")	
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf("- %v\n",encounter.Pokemon.Name)
	}
	return nil
}
