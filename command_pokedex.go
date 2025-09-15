package main

import (
	"fmt"
)

func commandPokedex(c *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for name := range c.caughtPokemon {
		fmt.Println("  - ", name)
	}
	return nil
}
