package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.prevLocationAreaUrl = resp.Previous
	cfg.nextLocationAreaUrl = resp.Next

	return nil
}

func callbackMapB(cfg *config) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.prevLocationAreaUrl = resp.Previous
	cfg.nextLocationAreaUrl = resp.Next

	return nil
}
