package main

import (
	"fmt"
	"errors"
	"math/rand"
)
func commandCatch(c *config, args ...string) error {
	if args == nil || len(args) == 0 {
		return errors.New("No Pokemon specified to catch...")
	}

	pokemon, err := c.pokeapiClient.PokemonInfo(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	//fmt.Println(pokemon.BaseExp)
	catchRate := 1 - 0.003*float64(pokemon.BaseExp)
	roll := rand.Float64()

	//fmt.Printf("Catch Rate: %.3v, Roll: %.3v\n", catchRate, roll)
	if roll < catchRate {
		fmt.Printf("%v was captured!!\n", pokemon.Name)
		c.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v broke out of the pokeball!\n", pokemon.Name)
	}

	return nil
}
