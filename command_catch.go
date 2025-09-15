package main

import (
	"fmt"
	"errors"
)
func commandCatch(c *config, args ...string) error {
	if args == nil || len(args) == 0 {
		return errors.New("No Pokemon specified to catch...")
	}
	fmt.Printf("Throwing a Pokeball at %v...", args[0])

	return nil
}
