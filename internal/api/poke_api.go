package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(pageURL *string) (RespShallowLocations, error) {
	requestURL := baseURL + "/location-area"
	if pageURL != nil {
		requestURL = *pageURL
	}
	res, err := c.httpClient.Get(requestURL)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	locationsData := RespShallowLocations{}

	if err := json.NewDecoder(res.Body).Decode(&locationsData); err != nil {
		return RespShallowLocations{}, err
	}
	return locationsData, nil
}

func (c *Client) GetLocationAreasAlt(pageURL *string) (RespShallowLocations, error) {
	requestURL := baseURL + "/location-area"
	if pageURL != nil {
		requestURL = *pageURL
	}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsData := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsData)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsData, nil
}
