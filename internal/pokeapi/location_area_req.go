package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

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

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	LocAreaResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &LocAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return LocAreaResp, nil
}
