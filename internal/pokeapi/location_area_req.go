package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		// ** CACHE READ **
		var locationAreas LocationAreasResp
		err := json.Unmarshal([]byte(dat), &locationAreas)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	var locationAreas LocationAreasResp
	err = json.Unmarshal([]byte(dat), &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// ** CACHE SAVE **
	c.cache.Add(fullURL, dat)
	return locationAreas, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationAreaResp, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		// ** CACHE READ **
		var locationArea LocationAreaResp
		err := json.Unmarshal([]byte(dat), &locationArea)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	var locationArea LocationAreaResp
	err = json.Unmarshal([]byte(dat), &locationArea)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// ** CACHE SAVE **
	c.cache.Add(fullURL, dat)
	return locationArea, nil
}
