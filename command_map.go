package main

import (
	"fmt"
)

func commandMapf(c *config) error {
	data, err := c.pokeapiClient.ListLocations(c.Next)
	if err != nil {
		return err//fmt.Errorf("Error getting pokeapi locations: %w", err)
	}
	c.Previous = data.Previous
	c.Next = data.Next

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	
	data, err := c.pokeapiClient.ListLocations(c.Previous)
	if err != nil {
		return err
	}
	c.Next = data.Next
	c.Previous = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
