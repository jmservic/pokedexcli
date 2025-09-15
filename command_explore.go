package main

import (
	"fmt"
	"errors"
)

func commandExplore( c *config, args ...string) error {
	if args == nil || len(args) == 0 {
		return errors.New("No Location name or id specified.")
	}
	fmt.Printf("Exploring %v...\n", args[0])
	data, err := c.pokeapiClient.ExploreLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon")	
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf("- %v\n",encounter.Pokemon.Name)
	}
	return nil
}
