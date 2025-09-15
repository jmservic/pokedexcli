package main

import (
	"fmt"
	"errors"

)

func commandInspect(c *config, args ...string) error {
	if args == nil || len(args) == 0 {
		return errors.New("No pokemon specified to inspect.")
	}
	pokemon, ok := c.caughtPokemon[args[0]]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, statInfo := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", statInfo.Stat.Name, statInfo.BaseStat)	
	}

	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  - %v\n", typeInfo.Type.Name)
	}

	return nil
}
