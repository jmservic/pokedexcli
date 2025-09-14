package main

import (
	"fmt"
	"errors"
)
func commandCatch(c *config) error {
	if c.Args == nil || len(c.Args) == 0 {
		return errors.New("No Pokemon specified to catch...")
	}
	fmt.Printf("Throwing a Pokeball at %v...", c.Args[0])

	return nil
}
