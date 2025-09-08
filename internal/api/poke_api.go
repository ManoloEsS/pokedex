package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient(timeout int) *Client {
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	return &Client{
		httpClient: client,
	}
}

func (c *Client) GetLocationAreas(url *string) (RespShallowLocations, error) {
	requestURL := BaseUrl + "/location-area"
	if url != nil {
		requestURL = *url
	}
	res, err := c.httpClient.Get(requestURL)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	locationsData := RespShallowLocations{}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&locationsData); err != nil {
		return RespShallowLocations{}, err
	}
	return locationsData, nil
}
