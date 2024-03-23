package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonResp, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		// ** CACHE READ **
		var Pokemon PokemonResp
		err := json.Unmarshal([]byte(dat), &Pokemon)
		if err != nil {
			return PokemonResp{}, err
		}
		return Pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	var Pokemon PokemonResp
	err = json.Unmarshal([]byte(dat), &Pokemon)
	if err != nil {
		return PokemonResp{}, err
	}

	// ** CACHE SAVE **
	c.cache.Add(fullURL, dat)
	return Pokemon, nil
}
