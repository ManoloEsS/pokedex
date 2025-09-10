package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
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
	requestURL, err := url.JoinPath(baseURL, "location-area")
	if err != nil {
		return RespShallowLocations{}, err
	}
	if pageURL != nil {
		requestURL = *pageURL
	}

	if val, ok := c.cache.Get(requestURL); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
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

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(requestURL, data)
	return locationsResp, nil
}
