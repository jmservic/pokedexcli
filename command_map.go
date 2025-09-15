package main

import (
	"fmt"
)

func commandMapf(c *config, args ...string) error {
	data, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err//fmt.Errorf("Error getting pokeapi locations: %w", err)
	}
	c.prevLocationsURL = data.Previous
	c.nextLocationsURL = data.Next

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(c *config, args ...string) error {
	if c.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	
	data, err := c.pokeapiClient.ListLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}
	c.nextLocationsURL = data.Next
	c.prevLocationsURL = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
