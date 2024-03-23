package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.prevLocationAreaUrl = resp.Previous
	cfg.nextLocationAreaUrl = resp.Next

	return nil
}

func callbackMapB(cfg *config, args ...string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.prevLocationAreaUrl = resp.Previous
	cfg.nextLocationAreaUrl = resp.Next

	return nil
}
